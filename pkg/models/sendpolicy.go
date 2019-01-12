package models

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"time"
)

// use to control time to send fired alert to user or webhooks(use webhook and system wehhook)
type SendPolicy struct {
	//SendPolicyID              string    `gorm:"primary_key"`
	ResourceID                string    `gorm:"primary_key"`
	AlertRuleID               string    `gorm:"primary_key"`
	CurrentRepeatSendInterval uint32    `gorm:"type:int unsigned;not null;"`
	CumulateRepeatSendCount   uint32    `gorm:"type:int unsigned;not null;"`
	CurrentRepeatSendAt       time.Time `gorm:"not null;"`
	SilenceStartAt            time.Time `gorm:""`
	SilenceEndAt              time.Time `gorm:""`
	CreatedAt                 time.Time `gorm:"not null;"`
	UpdatedAt                 time.Time `gorm:"not null;"`
}

func CreateSendPolicy(sendPolicy *SendPolicy) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	// get fired sendPolicy by
	resID := sendPolicy.ResourceID
	ruleID := sendPolicy.AlertRuleID

	if resID == "" || ruleID == "" {
		return Error{Text: "resource id and rule id must be specified", Code: InvalidParam}
	}

	//sendPolicy.SendPolicyID = idutil.GetUuid36("send_policy-")

	err = db.Model(&SendPolicy{}).Create(sendPolicy).Error

	return Error{Text: err.Error(), Code: DBError}
}

func CreateOrUpdateSendPolicy(sendPolicy *SendPolicy) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	if err := db.Save(&sendPolicy).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	return err
}

func GetSendPolicy(sendPolicy *SendPolicy) (*SendPolicy, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	resID := sendPolicy.ResourceID
	ruleID := sendPolicy.AlertRuleID

	if resID == "" || ruleID == "" {
		return nil, Error{Text: "resource id and rule id must be specified", Code: DBError}
	}

	var policy SendPolicy
	if db.Where("resource_id=? AND alert_rule_id=?", resID, ruleID).First(&policy).RecordNotFound() {
		return nil, nil
	}

	return &policy, Error{Text: db.Error.Error(), Code: DBError}
}

func GetOrCreateSendPolicy(sp *SendPolicy) (*SendPolicy, error) {
	sendPolicy, err := GetSendPolicy(sp)
	if err != nil {
		return nil, err
	} else {
		if sendPolicy == nil {
			err := CreateSendPolicy(sp)

			if err != nil {
				return nil, err
			} else {
				return sp, nil
			}

		} else {
			return sendPolicy, nil
		}
	}
}
