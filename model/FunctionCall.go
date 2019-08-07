package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
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
	argumentValues, err := exp.FunctionArguments.EvaluateArguments()
	if err != nil {
		return reflect.ValueOf(nil), err
	} else {
		switch exp.FunctionName {
		case "now":
			return reflect.ValueOf(time.Now()), nil
		case "isNil":
			if len(argumentValues) != 1 {
				return reflect.ValueOf(nil), fmt.Errorf("isNil function requires 1 parameter")
			} else {
				return reflect.ValueOf(argumentValues[0].IsNil()), nil
			}
		case "log":
			if len(argumentValues) != 1 {
				return reflect.ValueOf(nil), fmt.Errorf("log function requires 1 string parameter")
			} else {
				return reflect.ValueOf(argumentValues[0].IsNil()), nil
			}
		default:
			return reflect.ValueOf(nil), fmt.Errorf("unrecognized function %s", exp.FunctionName)
		}
	}
}
