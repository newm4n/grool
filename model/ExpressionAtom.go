package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/pkg"
	"reflect"
)

type ExpressionAtom struct {
	ExpressionAtomLeft  *ExpressionAtom
	ExpressionAtomRight *ExpressionAtom
	MathOperator        MathOperator
	Variable            string
	Constant            *Constant
	FunctionCall        *FunctionCall
	MethodCall          *MethodCall
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (exp *ExpressionAtom) Evaluate() (reflect.Value, error) {
	if len(exp.Variable) > 0 {
		return exp.dataCtx.GetValue(exp.Variable)
	} else if exp.Constant != nil {
		return exp.Constant.Evaluate()
	} else if exp.FunctionCall != nil {
		return exp.FunctionCall.Evaluate()
	} else {
		lv, err := exp.ExpressionAtomLeft.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		rv, err := exp.ExpressionAtomRight.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		switch exp.MathOperator {
		case MathOperatorPlus:
			return pkg.ValueAdd(lv, rv)
		case MathOperatorMinus:
			return pkg.ValueSub(lv, rv)
		case MathOperatorMul:
			return pkg.ValueMul(lv, rv)
		case MathOperatorDiv:
			return pkg.ValueDiv(lv, rv)
		}
		return reflect.ValueOf(nil), errors.Errorf("math operation can only be applied to numerical data (eg. int, uit or float) or string")
	}
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

	if ins.MethodCall != nil {
		ins.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (expr *ExpressionAtom) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if expr.ExpressionAtomLeft == nil {
		expr.ExpressionAtomLeft = exprAtom
	} else if expr.ExpressionAtomRight == nil {
		expr.ExpressionAtomRight = exprAtom
	} else {
		return errors.Errorf("expression alredy set twice")
	}
	return nil
}

func (expr *ExpressionAtom) AcceptFunctionCall(funcCall *FunctionCall) error {
	if expr.FunctionCall != nil {
		return errors.Errorf("functioncall alredy set")
	}
	expr.FunctionCall = funcCall
	return nil
}

func (ins *ExpressionAtom) AcceptMethodCall(methodCall *MethodCall) error {
	if ins.MethodCall != nil {
		return errors.Errorf("method call alredy set")
	}
	ins.MethodCall = methodCall
	return nil
}

func (expr *ExpressionAtom) AcceptVariable(name string) error {
	if expr.Variable == "" {
		expr.Variable = name
		return nil
	} else {
		return errors.Errorf("variable already defined")
	}
}

func (expr *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if expr.Constant == nil {
		expr.Constant = cons
		return nil
	} else {
		return errors.Errorf("constant already defined")
	}
}
