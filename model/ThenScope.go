package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
)

type ThenScope struct {
	AssignExpressions *AssignExpressions
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *context.RuleContext
	dataCtx           *context.DataContext
}

func (then *ThenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	then.knowledgeContext = knowledgeContext
	then.ruleCtx = ruleCtx

	if then.AssignExpressions != nil {
		then.AssignExpressions.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (then *ThenScope) Execute() error {
	_, err := then.AssignExpressions.Evaluate()
	if err != nil {
		return errors.Trace(err)
	} else {
		return nil
	}
}
