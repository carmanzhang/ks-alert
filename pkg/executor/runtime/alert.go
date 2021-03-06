package runtime

import (
	"context"
	"errors"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/executor/metric"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/notification"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"github.com/golang/glog"
	log "log"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var olderTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

const (
	MaxKeepAliveReportInterval = 3
	TickerFrequence            = 8
)

// this chan is used to control corresponding goroutine
type RuntimeAlertConfig struct {
	SigCh            chan pb.Informer_Signal
	StatusCh         chan StatusType
	UpdatedAt        time.Time
	CreatedAt        time.Time
	err              error
	alertConfig      *models.AlertConfig
	uriPath          string
	queryParams      string
	resourceNames    map[string]string
	freq             []int32
	currentFreq      []int32
	ruleEnable       []bool
	firedAlerts      map[string]map[string]time.Time
	keepAliveCounter int
}

// goroutine status
type StatusType string

const (
	Alive  StatusType = "alive"
	Dead   StatusType = "dead"
	Unkonw StatusType = "unknow"
)

type AlertStatus struct {
	Status    StatusType
	timestamp time.Time
}

// key is alert config id
// value is alert runtime parameters included channels
type RuntimeAlertStatus struct {
	sync.RWMutex
	Map map[string]*RuntimeAlertConfig
}

var CachedRuntimeAlert = &RuntimeAlertStatus{
	Map: make(map[string]*RuntimeAlertConfig),
}

func ExecuteAlertConfig(ctx context.Context, msg *pb.Informer) Error {

	switch msg.Signal {
	case pb.Informer_CREATE:
		if _, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]; ok {

			return Error{Text: "the alert config is under executing", Code: RuntimeError, Where: Caller(0, true)}
		}
		// create alert by specifig alert config id within one goroutine
		var rtAlert = &RuntimeAlertConfig{
			SigCh:     make(chan pb.Informer_Signal),
			CreatedAt: time.Now(),
		}

		err := rtAlert.reload(msg.AlertConfigId)

		if err != nil {
			return Error{Text: err.Error(), Code: RuntimeError, Where: Caller(0, true)}
		}

		CachedRuntimeAlert.Lock()
		CachedRuntimeAlert.Map[msg.AlertConfigId] = rtAlert
		CachedRuntimeAlert.Unlock()
		glog.Infof("executing alert config %s", msg.AlertConfigId)
		go rtAlert.runAlert()

	case pb.Informer_TERMINATE:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return Error{Text: "alert config was not executor by executor", Code: RuntimeError, Where: Caller(0, true)}
		}

		runtimeAlert.SigCh <- pb.Informer_TERMINATE

		CachedRuntimeAlert.Lock()
		delete(CachedRuntimeAlert.Map, msg.AlertConfigId)
		CachedRuntimeAlert.Unlock()

	case pb.Informer_RELOAD:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return Error{Text: "alert config was not executor by executor", Code: RuntimeError, Where: Caller(0, true)}
		}

		runtimeAlert.SigCh <- pb.Informer_RELOAD

	default:
		runtimeAlert, ok := CachedRuntimeAlert.Map[msg.AlertConfigId]
		if !ok {
			return Error{Text: "alert config was not executor by executor", Code: RuntimeError, Where: Caller(0, true)}
		}
		runtimeAlert.StatusCh <- Alive
		status := <-runtimeAlert.StatusCh
		log.Printf("%s,%s", msg.AlertConfigId, status)
	}

	return Error{Text: "success", Code: Success, Where: Caller(0, true)}
}

func (rtAlert *RuntimeAlertConfig) runAlert() {
	d := time.Duration(time.Second * time.Duration(TickerFrequence))
	timer := time.NewTicker(d)
	defer timer.Stop()

	sigCh := rtAlert.SigCh
	statusCh := rtAlert.StatusCh
	for {
		select {
		case <-statusCh:
			statusCh <- Alive

		case sig := <-sigCh:
			if sig == pb.Informer_RELOAD {
				acID := rtAlert.alertConfig.AlertConfigID
				log.Println("runtime rtAlert was reload, alert_config_id is: ", acID)
				var err error
				err = rtAlert.reload(acID)

				if err != nil {
					glog.Errorln(err.Error())
					rtAlert.err = err
					return
				}

			} else if sig == pb.Informer_TERMINATE {
				log.Println("runtime rtAlert was terminated, alert_config_id is: ", rtAlert.alertConfig.AlertConfigID)
				return
			}

		case <-timer.C:
			counter := rtAlert.keepAliveCounter
			counter = (counter + 1) % MaxKeepAliveReportInterval
			rtAlert.keepAliveCounter = counter

			if counter == 0 {
				go func() {
					err := models.UpdateAlertConfigKeepAliveTime(rtAlert.alertConfig.AlertConfigID)
					if err != nil {
						glog.Errorln(err.Error())
					}
				}()
			}

			// TODO need to add exception catcher function
			log.Println("new evalted period: ", len(CachedRuntimeAlert.Map), runtime.NumGoroutine(), counter)
			log.Println("fired alert: ", jsonutil.Marshal(rtAlert.firedAlerts))
			alertConfig := rtAlert.alertConfig

			// 0. check this alert_config's hostid, whether is consistency with this node or not
			// this step is important, because when network became unreachable, alert system will launch another node(pod)
			// to take replace of this node(pod), it means the same alert configs will be executed in a newly created node
			// if both are not consistency, goroutine will be exist.
			hostID, err := models.GetAlertConfigBindingHost(alertConfig.AlertConfigID)
			if err != nil {
				glog.Errorln(err.Error())
			}

			if hostID != option.HostInfo {
				return
			}

			//1. rtAlert config is enable?
			enable := isAlertEnable(alertConfig.EnableStart, alertConfig.EnableEnd)
			if !enable {
				continue
			}

			// 2. is there any rules need to execute?
			evaluatedRuleIndx := getExecutingRuleIndx(rtAlert.freq, rtAlert.currentFreq, rtAlert.ruleEnable)
			if len(evaluatedRuleIndx) == 0 {
				continue
			}

			// 3. retrieve metrics from monitoring center using evaluating rules
			var ch = make(chan *metric.ResourceMetrics, 15)
			rtAlert.getResourceMetrics(evaluatedRuleIndx, ch)
			close(ch)

			// 4. decide whether a resource trigger a rule
			for metricByRule := range ch {
				rtAlert.evaluteAlertInPipeline(metricByRule)
			}
		}

	}
}

func (rtAlert *RuntimeAlertConfig) evaluteAlertInPipeline(metricByRule *metric.ResourceMetrics) {
	rules := rtAlert.alertConfig.AlertRuleGroup.AlertRules
	if metricByRule != nil {
		indx := metricByRule.RuleIndx
		rule := rules[indx]
		consCount := int(rule.ConsecutiveCount)
		resourceMetric := metricByRule.ResourceMetric
		resNameMap := rtAlert.resourceNames
		for resName := range resourceMetric {
			// timestamp and value
			tvs := resourceMetric[resName]
			ll := len(tvs)
			if ll < consCount {
				continue
			}

			tvs = tvs[ll-consCount:]
			ts := tvs[len(tvs)-1].T
			lastEvalutedTime := time.Unix(int64(ts), 0)

			// 0. is alert fired?
			isFired := isFired(rule, tvs)

			// 1. update alert status (fired or recovered)
			ruleID := rule.AlertRuleID
			isRecovery := rtAlert.updateAlertFiredStatus(resName, ruleID, isFired)

			// 2. insert an record(fired alert record or recovery record)
			// insert alert recovery item into `alert_histories`
			// `tvs` will be send to user and insert into `alert_histories` if alert fired
			var ah *models.AlertHistory
			if isFired || isRecovery {
				ah = rtAlert.makeAlertHistoryItem(rule, resName, lastEvalutedTime, tvs, isRecovery)
				_, err := models.CreateAlertHistory(ah)
				if err != nil {
					glog.Errorln(err.Error())
					continue
				}
			}

			if !isFired {
				continue
			}

			// 3. get send policies (silence and repeat send policy) from db
			resID := resNameMap[resName]
			sendPolicy, err := models.GetOrCreateSendPolicy(&models.SendPolicy{AlertRuleID: rule.AlertRuleID, ResourceID: resID})
			if err != nil {
				glog.Errorln(err.Error())
				continue
			}

			//4. check it's the time for the fired alert is going to send
			needSend, isResendIntervalChanged := checkSendSatisfied(sendPolicy, rule)
			if !needSend {
				continue
			}

			//5. check repeat send policy satisfied, this policy may need to update
			if isResendIntervalChanged {
				sendPolicy.NextRepeatSendInterval = rule.InitRepeatSendInterval
			}
			currentRepeatSendInterval := sendPolicy.NextRepeatSendInterval

			updatedPolicy := updateSendPolicy(sendPolicy, rule)
			nextRepeatSendInterval := updatedPolicy.NextRepeatSendInterval
			cumulateRepeatSendCount := updatedPolicy.CumulateRepeatSendCount

			err = models.CreateOrUpdateSendPolicy(updatedPolicy)
			if err != nil {
				glog.Errorln(err.Error())
				continue
			}

			// 6. make notice
			sendNoticeAt := time.Now()
			notice := notification.Notice{
				ResourceName:          resName,
				Metrics:               &tvs,
				Rule:                  rule,
				TriggerTime:           lastEvalutedTime,
				CumulateReSendCount:   cumulateRepeatSendCount,
				SendNoticeAt:          sendNoticeAt,
				MaxReSendCount:        rule.MaxRepeatSendCount,
				CurrentReSendInterval: currentRepeatSendInterval,
			}

			if cumulateRepeatSendCount < rule.MaxRepeatSendCount {
				notice.NextReSendInterval = nextRepeatSendInterval
			}

			noticeStr := notice.MakeNotice(false)

			// 7. send notice
			sendStatus := notification.Sender{}.Send(rtAlert.alertConfig.ReceiverGroup.Receivers, noticeStr)

			// 8. update send status and send policy in `alert_history`
			ah.RequestNotificationStatus = sendStatus
			ah.NextRepeatSendInterval = nextRepeatSendInterval
			ah.CumulateRepeatSendCount = cumulateRepeatSendCount
			ah.NotificationSendAt = sendNoticeAt
			if !updatedPolicy.SilenceStartAt.Equal(olderTime) {
				ah.SilenceStartAt = updatedPolicy.SilenceStartAt
			}

			if !updatedPolicy.SilenceEndAt.Equal(olderTime) {
				ah.SilenceEndAt = updatedPolicy.SilenceEndAt
			}

			err = models.UpdateAlertHistory(ah)
			if err != nil {
				glog.Error(err.Error())
				continue
			}
		}
	}
}

// checking whether fired alert has been recovery or not
// remove alert fired alert info from `firedAlerts` if exist
func (rtAlert *RuntimeAlertConfig) updateAlertFiredStatus(resName string, ruleID string, isFired bool) bool {
	firedAlerts := rtAlert.firedAlerts
	if firedAlerts == nil {
		firedAlerts = make(map[string]map[string]time.Time)
	}

	isRecovery := false

	if isFired {
		// add fired alert info into `firedAlerts` if not exist
		if firedRules, ok := firedAlerts[resName]; ok && firedRules != nil {
			if _, ok := firedRules[ruleID]; !ok {
				firedRules[ruleID] = time.Now()
			}

		} else {
			firedRules = make(map[string]time.Time)
			firedRules[ruleID] = time.Now()
			firedAlerts[resName] = firedRules
		}
	} else {
		if firedRules, ok := firedAlerts[resName]; ok && firedRules != nil {
			if _, ok := firedRules[ruleID]; ok {
				delete(firedRules, ruleID)
				isRecovery = true
				//log.Printlnln("recovery alert", rules[indx].MetricName, resName)
			}

			if len(firedRules) == 0 {
				delete(firedAlerts, resName)
			}
		}
	}

	return isRecovery
}

func isFired(rule *models.AlertRule, tvs []metric.TV) bool {
	condition := rule.ConditionType
	threshold := rule.Threshold
	f := true
	for i := 0; i < len(tvs); i++ {
		v, err := strconv.ParseFloat(tvs[i].V, 32)
		if err != nil {
		}
		switch condition {
		case ">=":
			f = f && (v >= float64(threshold))
		case ">":
			f = f && (v > float64(threshold))
		case "<=":
			f = f && (v <= float64(threshold))
		case "<":
			f = f && (v < float64(threshold))
		}

		if !f {
			break
		}
	}
	return f
}

func checkSendSatisfied(sendPolicy *models.SendPolicy, rule *models.AlertRule) (bool, bool) {
	// check silence policy
	//silenceStart := sendPolicy.SilenceStartAt
	var needResend = false
	var isResendIntervalChanged = false
	silenceEnd := sendPolicy.SilenceEndAt
	now := time.Now()
	if silenceEnd.Before(now) {
		// silence policy lose effectiveness，
		// does not take any effect, need to change the SilenceStart and SilenceEnd
		//pasedTime, _ := time.Parse("0001-01-01T00:00:00Z", "")
		// check repeat send policy
		maxSendCount := rule.MaxRepeatSendCount

		currReSendTime := sendPolicy.CurrentRepeatSendAt
		alreaySendCount := sendPolicy.CumulateRepeatSendCount
		nextReSendInterval := sendPolicy.NextRepeatSendInterval

		if alreaySendCount < maxSendCount {
			if alreaySendCount == 0 {
				needResend = true
			} else {

				// maybe rule was updated, but the sendPolicy does not know
				oldInitReSendInterval := sendPolicy.InitRepeatSendInterval
				initResendInterval := rule.InitRepeatSendInterval
				if oldInitReSendInterval != 0 && oldInitReSendInterval != initResendInterval {
					nextReSendInterval = initResendInterval
					isResendIntervalChanged = true
				}

				d := time.Minute * time.Duration(nextReSendInterval)

				nextReSendTime := currReSendTime.Add(d)
				if !now.Before(nextReSendTime) {
					needResend = true
				} else {
					// dees not reach next repeat send interval
					log.Println("dees not reach next repeat send interval", now, nextReSendTime, nextReSendTime.Sub(now))
				}
			}
		} else {
			// exceed maximum repeat send count, this fired alert will be inhibited
			log.Println("exceed maximum repeat send count, this fired alert will be inhibited")
		}
	} else {
		// still in silence period
		log.Println("still in silence period", sendPolicy.SilenceStartAt, sendPolicy.SilenceEndAt)
	}

	return needResend, isResendIntervalChanged
}

func updateSendPolicy(sendPolicy *models.SendPolicy, rule *models.AlertRule) *models.SendPolicy {

	var newInterval uint32 = 0

	initRepeatSendInterval := rule.InitRepeatSendInterval

	if sendPolicy.CumulateRepeatSendCount == 0 {
		newInterval = initRepeatSendInterval
	} else {
		// update next repeat interval, because this fired alert will be send to users
		// `NextRepeatSendInterval` is the current repeat interval
		currentResendInterval := sendPolicy.NextRepeatSendInterval
		if currentResendInterval == 0 {
			newInterval = initRepeatSendInterval
		} else {
			newInterval = NextReSendInterval(rule.RepeatSendType, currentResendInterval, initRepeatSendInterval)
		}
	}

	sendPolicy.InitRepeatSendInterval = initRepeatSendInterval
	sendPolicy.NextRepeatSendInterval = newInterval
	sendPolicy.CumulateRepeatSendCount += 1
	sendPolicy.CurrentRepeatSendAt = time.Now()

	return sendPolicy
}

func NextReSendInterval(sendType int32, currRepeatSendInterval, initSendInterval uint32) uint32 {
	// 	0: "Fixed",
	//	1: "Exponential",
	var nextReSendInterval uint32 = 0

	if currRepeatSendInterval == 0 {
		nextReSendInterval = initSendInterval
	} else {
		if sendType == 1 {
			nextReSendInterval = currRepeatSendInterval * 2
		} else {
			nextReSendInterval = currRepeatSendInterval
		}
	}

	return nextReSendInterval
}

// alert rules which are need to execute
func getExecutingRuleIndx(freq, currentFreq []int32, ruleEnable []bool) []int {
	l := len(freq)
	var evaluatedRuleIndx []int
	for i := 0; i < l; i++ {
		if !ruleEnable[i] {
			continue
		}
		currentFreq[i] = (currentFreq[i] + 1) % freq[i]
		if currentFreq[i] == 0 {
			// means this rule whill be evaluated in turn
			evaluatedRuleIndx = append(evaluatedRuleIndx, i)
		}
	}
	log.Println("freq: ", currentFreq, freq)
	return evaluatedRuleIndx
}

func (rtAlert *RuntimeAlertConfig) getResourceMetrics(evaluatedRuleIndx []int, ch chan *metric.ResourceMetrics) {

	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	resourceNames := rtAlert.resourceNames
	uriPath := rtAlert.uriPath
	queryParams := rtAlert.queryParams

	rules := rtAlert.alertConfig.AlertRuleGroup.AlertRules

	var resNameArray []string
	for n := range resourceNames {
		resNameArray = append(resNameArray, n)
	}

	for i := 0; i < len(evaluatedRuleIndx); i++ {
		j := evaluatedRuleIndx[i]
		rule := rules[j]
		wg.Add(1)
		go func(rule *models.AlertRule) {
			for {
				select {
				case <-ctx.Done():
					glog.Warningf("pull metrics was canceled, metric name: %s, resource name: %v", rule.MetricName, resNameArray)
					wg.Done()
					return
				default:
					metricName := rule.MetricName
					// period unit is Minute
					stepInMinute := rule.Period
					consCount := rule.ConsecutiveCount
					if consCount <= 0 || consCount > metric.MaxConsecutiveCount {
						consCount = metric.ConsecutiveCount
					}
					if stepInMinute <= 0 || stepInMinute > metric.MaxStep {
						stepInMinute = metric.Step
					}
					endTime := time.Now().Truncate(time.Minute).Unix()
					startTime := endTime - int64((consCount+3)*stepInMinute*60)
					metricStr := client.SendMonitoringRequest(uriPath, queryParams, resNameArray, metricName, startTime, endTime, stepInMinute)
					resourceMetrics := metric.GetResourceTimeSeriesMetric(metricStr, metricName, startTime, endTime)
					if resourceMetrics != nil {
						resourceMetrics.RuleIndx = j
						log.Println("pull metrics: ", resourceMetrics)
						ch <- resourceMetrics
					}
					wg.Done()
					return
				}
			}
		}(rule)
	}
	wg.Wait()
}

func isAlertEnable(start time.Time, end time.Time) bool {
	now := time.Now()

	if start.Before(now) && now.Before(end) {
		return true
	}

	return false
}

// reload alert config
func (rtAlert *RuntimeAlertConfig) reload(acID string) error {
	// get alert config by id from DB
	alertConfig, err := models.GetAlertConfig(&models.AlertConfig{AlertConfigID: acID})

	if err != nil {
		return err
	}

	alertRules := alertConfig.AlertRuleGroup
	if alertRules == nil || len(alertRules.AlertRules) == 0 {
		return errors.New("at least one alert rule must be specified")
	}

	receivers := alertConfig.ReceiverGroup
	if receivers == nil || len(*receivers.Receivers) == 0 {
		return errors.New("at least one receiver must be specified")
	}

	resources := alertConfig.ResourceGroup
	if resources == nil || len(resources.Resources) == 0 {
		return errors.New("at least one resource must be specified")
	}

	uriPath, queryParams, resNames, err := GetResourcesSpec(alertConfig.ResourceGroup)

	if err != nil {
		return err
	}

	rtAlert.alertConfig = alertConfig
	rtAlert.resourceNames = resNames
	rtAlert.uriPath = uriPath
	rtAlert.queryParams = queryParams

	// TODO need to keep consistency with status before this reload
	l := len(alertRules.AlertRules)
	var freq = make([]int32, l)
	var currentFreq = make([]int32, l)
	var ruleEnable = make([]bool, l)

	for i := 0; i < l; i++ {
		rule := alertRules.AlertRules[i]
		freq[i] = rule.Period
		ruleEnable[i] = rule.Enable
	}

	rtAlert.freq = freq
	rtAlert.currentFreq = currentFreq
	rtAlert.ruleEnable = ruleEnable

	firedAlerts := rtAlert.firedAlerts
	if firedAlerts == nil {
		rtAlert.firedAlerts = make(map[string]map[string]time.Time)
	} else {
		removeOldRulesAndResources(resources, firedAlerts, alertRules)
	}
	rtAlert.UpdatedAt = time.Now()
	// TODO delete useless item in `send_policies`
	return nil
}

// is's necessary to clean up `firedAlerts`, for the reason reloading may discard old rules and add new rules
// `firedAlerts`  saved the fired alert triggered by an existing rule on a existing resource
// `firedAlerts` field was arranged by map[resource_name](map[alert_rule_id]2019.1.1-00:00:00)
// remove old resources
func removeOldRulesAndResources(resources *models.ResourceGroup, firedAlerts map[string]map[string]time.Time, alertRules *models.AlertRuleGroup) {
	l := len(resources.Resources)
	var resMap = make(map[string]string)

	for i := 0; i < l; i++ {
		rs := resources.Resources[i]
		resMap[rs.ResourceName] = ""
	}

	for k := range firedAlerts {
		if _, ok := resMap[k]; !ok {
			delete(firedAlerts, k)
		}
	}

	// 1. remove old rules
	l = len(alertRules.AlertRules)
	for k := range firedAlerts {
		currRuleMap := firedAlerts[k]
		for i := 0; i < l; i++ {
			ruleID := alertRules.AlertRules[i].AlertRuleID
			if _, ok := currRuleMap[ruleID]; !ok {
				delete(firedAlerts[k], ruleID)
			}
		}
	}
}

// get resources uriPath and resource names
func GetResourcesSpec(resourceGroup *models.ResourceGroup) (string, string, map[string]string, error) {

	var uriTmpl models.ResourceUriTmpl
	jsonutil.Unmarshal(resourceGroup.URIParams, &uriTmpl)

	pathParams := uriTmpl.PathParams
	queryParams := uriTmpl.QueryParams
	// find uriTmpls by resource type id
	resourceType, err := models.GetResourceType(&models.ResourceType{ResourceTypeID: resourceGroup.ResourceTypeID})
	if err != nil {
		return "", "", nil, err
	}

	monitorHost := strings.Trim(resourceType.MonitorCenterHost, " ")
	monitorPort := resourceType.MonitorCenterPort
	if monitorHost == "" || monitorPort == 0 {
		// get monitoring host:port by product id
		prod, err := models.GetProduct(&models.Product{ProductID: resourceType.ProductID})
		if err != nil {
			return "", "", nil, err
		}

		monitorHost = strings.Trim(prod.MonitorCenterHost, " ")
		monitorPort = prod.MonitorCenterPort
	}

	if monitorHost == "" || monitorPort == 0 {
		return "", "", nil, errors.New("monitoring center must be specified")
	}

	var uriTmpls models.ResourceUriTmpls
	jsonutil.Unmarshal(resourceType.ResourceURITmpls, &uriTmpls)

	l := len(uriTmpls.ResourceUriTmpl)

	b := false

	var urlTmpl string
	var storedQueryParams string
	for i := 0; i < l; i++ {
		u := uriTmpls.ResourceUriTmpl[i]
		// does pathParams match
		storedPathParams := u.PathParams
		if IsMatch(pathParams, storedPathParams) {
			urlTmpl = u.UriTmpl
			storedQueryParams = u.QueryParams
			b = true
			break
		}
	}

	if !b {
		return "", "", nil, errors.New("resource uriPath parameters dose not match any existing resource uriPath template")
	}

	if queryParams == "" {
		queryParams = storedQueryParams
	}

	uriPath, err := AssembleURLPrefix(monitorHost, monitorPort, urlTmpl, pathParams)

	if err != nil {
		return "", "", nil, err
	}

	resources := resourceGroup.Resources
	l = len(resources)
	var resNames = make(map[string]string, l)

	for i := 0; i < l; i++ {
		if resources[i] != nil {
			resNames[resources[i].ResourceName] = resources[i].ResourceID
		}
	}

	return uriPath, queryParams, resNames, nil
}

func AssembleURLPrefix(host string, port int32, uriTmpl string, params map[string]string) (string, error) {
	r, err := regexp.Compile(`\{\w+\}`)
	if err != nil {
		// compile error
		return "", err
	}

	uri := r.ReplaceAllStringFunc(uriTmpl, func(s string) string {
		s = strings.Trim(s, "{")
		s = strings.Trim(s, "}")
		return params[s]
	})

	uri = fmt.Sprintf("%s:%d%s", host, port, uri)
	return uri, nil
}

func IsMatch(p map[string]string, q map[string]string) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {

		if len(p) == len(q) {
			for k := range p {
				if _, ok := q[k]; !ok {
					return false
				}
			}

		} else {
			return false
		}

	} else {
		return false
	}

	return true
}

func (rtAlert *RuntimeAlertConfig) makeAlertHistoryItem(
	firedRule *models.AlertRule, resName string,
	lastEvalutedTime time.Time, tvs []metric.TV, isRecovery bool) *models.AlertHistory {
	ac := rtAlert.alertConfig
	ruleGroup := ac.AlertRuleGroup
	resGroup := ac.ResourceGroup
	recvGroup := ac.ReceiverGroup

	ah := models.AlertHistory{}
	ah.AlertConfigID = ac.AlertConfigID
	//ah.ProductID =
	ah.ReceiverGroupID = ac.ReceiverGroupID
	ah.ReceiverGroup = fmt.Sprintf("%v", recvGroup)
	ah.ReceiverGroupName = recvGroup.ReceiverGroupName

	ah.ResourceGroupID = ac.ResourceGroupID
	ah.ResourceGroupName = resGroup.ResourceGroupName
	ah.AlertedResource = resName

	ah.AlertRuleGroupID = ruleGroup.AlertRuleGroupID
	ah.TriggerMetricName = firedRule.MetricName
	ah.AlertRuleGroupName = ruleGroup.AlertRuleGroupName

	ah.SeverityID = ac.SeverityID
	ah.SeverityCh = ac.SeverityCh

	ah.RepeatSendType = uint32(firedRule.RepeatSendType)
	ah.InitRepeatSendInterval = firedRule.InitRepeatSendInterval
	ah.MaxRepeatSendCount = firedRule.MaxRepeatSendCount

	ah.CreatedAt = time.Now()
	if isRecovery {
		ah.AlertRecoveryAt = lastEvalutedTime
	} else {
		ah.AlertFiredAt = lastEvalutedTime
	}

	ah.MetricData = jsonutil.Marshal(tvs)
	return &ah
}
