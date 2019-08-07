package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/pkg"
	"log"
	"reflect"
	"time"
)

type FunctionCall struct {
	FunctionName      string
	FunctionArguments *FunctionArgument
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *context.RuleContext
	dataCtx           *context.DataContext
}

func (ins *FunctionCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.FunctionArguments != nil {
		ins.FunctionArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (exp *FunctionCall) Evaluate() (reflect.Value, error) {
	var argumentValues []reflect.Value
	if exp.FunctionArguments == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := exp.FunctionArguments.EvaluateArguments()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		argumentValues = av
	}
	switch exp.FunctionName {
	case "now":
		return reflect.ValueOf(time.Now()), nil
	case "isNil":
		if len(argumentValues) != 1 {
			return reflect.ValueOf(nil), errors.Errorf("isNil function requires 1 parameter")
		} else {
			if argumentValues[0].Kind() == reflect.Struct {
				z0 := reflect.Zero(argumentValues[0].Type())
				return reflect.ValueOf(pkg.ValueToInterface(z0) == pkg.ValueToInterface(argumentValues[0])), nil
			} else if argumentValues[0].Kind() == reflect.Ptr {
				return reflect.ValueOf(argumentValues[0].IsNil()), nil
			} else {
				return reflect.ValueOf(false), nil
			}
		}
	case "log":
		if len(argumentValues) != 1 {
			return reflect.ValueOf(nil), errors.Errorf("log function requires 1 string parameter")
		} else {
			stringVal := argumentValues[0].String()
			log.Println(stringVal)
			return reflect.ValueOf(nil), nil
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("unrecognized function %s", exp.FunctionName)
	}
}
