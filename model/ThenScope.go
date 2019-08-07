package model

import "github.com/newm4n/grool/context"

type ThenScope struct {
	AssignExpressions *AssignExpressions
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *context.RuleContext
	dataCtx           *context.DataContext
}

func (ins *ThenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx

	if ins.AssignExpressions != nil {
		ins.AssignExpressions.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ts *ThenScope) Execute() error {
	_, err := ts.AssignExpressions.Evaluate()
	if err != nil {
		return err
	} else {
		return nil
	}
}
