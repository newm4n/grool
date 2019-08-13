package model

import (
	"github.com/newm4n/grool/context"
	"reflect"
)

type MethodCall struct {
	MethodName       string
	MethodArguments  *FunctionArgument
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *MethodCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.MethodArguments != nil {
		ins.MethodArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ins *MethodCall) AcceptFunctionArgument(funcArg *FunctionArgument) error {
	ins.MethodArguments = funcArg
	return nil
}

func (exp *MethodCall) Evaluate() (reflect.Value, error) {
	// todo finish thiss
	return reflect.ValueOf(nil), nil
}
