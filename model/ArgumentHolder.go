package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
	"reflect"
)

type ArgumentHolder struct {
	Constant         *Constant
	Variable         string
	FunctionCall     *FunctionCall
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *ArgumentHolder) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.Constant != nil {
		ins.Constant.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.FunctionCall != nil {
		ins.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.Expression != nil {
		ins.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ins *ArgumentHolder) Evaluate() (reflect.Value, error) {
	if len(ins.Variable) > 0 {
		return ins.dataCtx.GetValue(ins.Variable)
	} else if ins.Constant != nil {
		return ins.Constant.Evaluate()
	} else if ins.FunctionCall != nil {
		return ins.FunctionCall.Evaluate()
	} else if ins.Expression != nil {
		return ins.Expression.Evaluate()
	} else {
		return reflect.ValueOf(nil), fmt.Errorf("argument holder stores no value")
	}
}
