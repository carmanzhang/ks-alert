package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/pkg/errors"
	"reflect"
)

type TP string

const (
	AlertConfig   TP = "AlertConfig"
	RuleGroup     TP = "RuleGroup"
	ReceiverGroup TP = "ReceiverGroup"
	ResourceGroup TP = "ResourceGroup"
	Severity      TP = "Severity"
)

const (
	MethodCreate = "Create"
	MethodUpdate = "Update"
	MethodGet    = "Get"
	MethodDelete = "Delete"
)

func DoTransactionAction(v interface{}, tp TP, method string) (interface{}, error) {

	// transaction begin
	db, e := dbutil.DBClient()
	if e != nil {
		return nil, e
	}

	tx := db.Begin()

	var res interface{}
	var err error

	if tp == AlertConfig {
		// create alert config itself include config name and description...
	}

	if tp == AlertConfig || tp == RuleGroup {
		// create rule group
	}

	if tp == AlertConfig || tp == ReceiverGroup {
		// receiver group
		res, err = CallReflect(models.ReceiverGroup{}, method, tx, v)
	}

	if tp == AlertConfig || tp == ResourceGroup {
		// create resource group
	}

	// TODO need to validate closing db connection
	// transaction end
	tx.Commit()

	return res, err
}

func CallReflect(any interface{}, method string, args ...interface{}) (interface{}, error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	v := reflect.ValueOf(any).MethodByName(method)

	if v.String() == "<invalid Value>" {
		return nil, errors.New("invalid Value")
	}

	values := v.Call(inputs)

	if len(values) == 1 {
		return values[0].Interface(), nil

	} else if len(values) == 2 {

		if values[1].Interface() == nil {
			return values[0].Interface(), nil
		}

		e := values[1].Interface().(error)
		return values[0].Interface(), e
	}

	return nil, nil
}
