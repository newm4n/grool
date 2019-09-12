package model

import "github.com/newm4n/grool/context"

// GraphInitializer defines all graph that can be initalized with context.
type GraphInitializer interface {
	Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext)
}
