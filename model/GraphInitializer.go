package model

import "github.com/newm4n/grool/context"

type GraphInitializer interface {
	Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext)
}
