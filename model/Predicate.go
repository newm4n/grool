package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
)

type Predicate struct {
	ExpressionAtomLeft  *ExpressionAtom
	ExpressionAtomRight *ExpressionAtom
	ComparisonOperator  ComparisonOperator
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

func (ins *Predicate) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.ExpressionAtomLeft != nil {
		ins.ExpressionAtomLeft.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.ExpressionAtomRight != nil {
		ins.ExpressionAtomRight.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (expr *Predicate) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if expr.ExpressionAtomLeft != nil {
		if expr.ExpressionAtomRight != nil {
			return fmt.Errorf("expression alredy set twice")
		}
		expr.ExpressionAtomRight = exprAtom
	}
	expr.ExpressionAtomLeft = exprAtom
	return nil
}
