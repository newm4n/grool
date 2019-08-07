package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
	"reflect"
)

type AssignExpression struct {
	Assignment       *Assignment
	FunctionCall     *FunctionCall
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *AssignExpression) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.Assignment != nil {
		ins.Assignment.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ins.FunctionCall != nil {
		ins.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ins *AssignExpression) AcceptFunctionCall(funcCall *FunctionCall) error {
	ins.FunctionCall = funcCall
	return nil
}

func (ins *AssignExpression) Evaluate() (reflect.Value, error) {
	if ins.Assignment != nil {
		return ins.Assignment.Evaluate()
	} else if ins.FunctionCall != nil {
		return ins.FunctionCall.Evaluate()
	} else {
		return reflect.ValueOf(nil), fmt.Errorf("no assignment or function call to evaluate")
	}
}
