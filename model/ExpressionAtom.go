package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
)

type ExpressionAtom struct {
	ExpressionAtomLeft  *ExpressionAtom
	ExpressionAtomRight *ExpressionAtom
	MathOperator        MathOperator
	Variable            string
	Constant            *Constant
	FunctionCall        *FunctionCall
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

func (ins *ExpressionAtom) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.ExpressionAtomLeft != nil {
		ins.ExpressionAtomLeft.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ins.ExpressionAtomRight != nil {
		ins.ExpressionAtomRight.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ins.Constant != nil {
		ins.Constant.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ins.FunctionCall != nil {
		ins.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (expr *ExpressionAtom) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if expr.ExpressionAtomLeft != nil {
		if expr.ExpressionAtomRight != nil {
			return fmt.Errorf("expression alredy set twice")
		}
		expr.ExpressionAtomRight = exprAtom
	}
	expr.ExpressionAtomLeft = exprAtom
	return nil
}

func (expr *ExpressionAtom) AcceptFunctionCall(funcCall *FunctionCall) error {
	if expr.FunctionCall != nil {
		return fmt.Errorf("functioncall alredy set")
	}
	expr.FunctionCall = funcCall
	return nil
}

func (expr *ExpressionAtom) AcceptVariable(name string) error {
	if expr.Variable == "" {
		expr.Variable = name
		return nil
	} else {
		return fmt.Errorf("variable already defined")
	}
}

func (expr *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if expr.Constant == nil {
		expr.Constant = cons
		return nil
	} else {
		return fmt.Errorf("constant already defined")
	}
}
