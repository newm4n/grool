package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

type AssignExpressions struct {
	ExpressionList   []*AssignExpression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *AssignExpressions) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.ExpressionList != nil {
		for _, val := range ins.ExpressionList {
			val.Initialize(knowledgeContext, ruleCtx, dataCtx)
		}
	}
}

func (ins *AssignExpressions) Evaluate() (reflect.Value, error) {
	for _, v := range ins.ExpressionList {
		_, err := v.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
	}
	return reflect.ValueOf(nil), nil
}
