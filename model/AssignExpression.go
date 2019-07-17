package model

import "github.com/newm4n/grool/context"

type AssignExpression struct {
	Assignment       *Assignment
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
}
