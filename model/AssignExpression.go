package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

type AssignExpression struct {
	Assignment       *Assignment
	FunctionCall     *FunctionCall
	MethodCall       *MethodCall
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ae *AssignExpression) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ae.knowledgeContext = knowledgeContext
	ae.ruleCtx = ruleCtx
	ae.dataCtx = dataCtx

	if ae.Assignment != nil {
		ae.Assignment.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ae.FunctionCall != nil {
		ae.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ae.MethodCall != nil {
		ae.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ae *AssignExpression) AcceptFunctionCall(funcCall *FunctionCall) error {
	ae.FunctionCall = funcCall
	return nil
}

func (ae *AssignExpression) AcceptMethodCall(methodCall *MethodCall) error {
	ae.MethodCall = methodCall
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ae *AssignExpression) Evaluate() (reflect.Value, error) {
	if ae.Assignment != nil {
		return ae.Assignment.Evaluate()
	} else if ae.FunctionCall != nil {
		return ae.FunctionCall.Evaluate()
	} else {
		return reflect.ValueOf(nil), errors.Errorf("no assignment or function call to evaluate")
	}
}
