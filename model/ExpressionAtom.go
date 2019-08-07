package model

import (
	"fmt"
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
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

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
			return reflect.ValueOf(nil), err
		}
		rv, err := exp.ExpressionAtomRight.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		if lv.Kind() == reflect.String && rv.Kind() == reflect.String && exp.MathOperator == MathOperatorPlus {
			return reflect.ValueOf(fmt.Sprintf("%s%s", lv.String(), rv.String())), nil
		} else {
			bkl := pkg.GetBaseKind(lv)
			bkr := pkg.GetBaseKind(rv)
			if (bkl == reflect.Int64 || bkl == reflect.Uint64 || bkl == reflect.Float64) &&
				(bkr == reflect.Int64 || bkr == reflect.Uint64 || bkr == reflect.Float64) {
				if bkl == reflect.Int64 && bkr == reflect.Int64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Int() + rv.Int()), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Int() - rv.Int()), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Int() * rv.Int()), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Int() / rv.Int()), nil
					}
				} else if bkl == reflect.Int64 && bkr == reflect.Uint64 {
					if rv.OverflowUint(rv.Uint()) {
						return reflect.ValueOf(nil), fmt.Errorf("the integer value %d is overflowing int64 type", rv.Uint())
					}
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Int() + int64(rv.Uint())), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Int() - int64(rv.Uint())), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Int() * int64(rv.Uint())), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Int() / int64(rv.Uint())), nil
					}
				} else if bkl == reflect.Int64 && bkr == reflect.Float64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(float64(lv.Int()) + rv.Float()), nil
					case MathOperatorMinus:
						return reflect.ValueOf(float64(lv.Int()) - rv.Float()), nil
					case MathOperatorMul:
						return reflect.ValueOf(float64(lv.Int()) * rv.Float()), nil
					case MathOperatorDiv:
						return reflect.ValueOf(float64(lv.Int()) / rv.Float()), nil
					}
				} else if bkl == reflect.Uint64 && bkr == reflect.Int64 {
					if lv.OverflowInt(rv.Int()) {
						return reflect.ValueOf(nil), fmt.Errorf("the integer value %d is overflowing int64 type", rv.Uint())
					}
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Uint() + uint64(rv.Int())), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Uint() - uint64(rv.Int())), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Uint() * uint64(rv.Int())), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Uint() / uint64(rv.Int())), nil
					}
				} else if bkl == reflect.Uint64 && bkr == reflect.Uint64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Uint() + rv.Uint()), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Uint() - rv.Uint()), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Uint() * rv.Uint()), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Uint() / rv.Uint()), nil
					}
				} else if bkl == reflect.Uint64 && bkr == reflect.Float64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(float64(lv.Uint()) + rv.Float()), nil
					case MathOperatorMinus:
						return reflect.ValueOf(float64(lv.Uint()) - rv.Float()), nil
					case MathOperatorMul:
						return reflect.ValueOf(float64(lv.Uint()) * rv.Float()), nil
					case MathOperatorDiv:
						return reflect.ValueOf(float64(lv.Uint()) / rv.Float()), nil
					}
				} else if bkl == reflect.Float64 && bkr == reflect.Int64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Float() + float64(rv.Int())), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Float() - float64(rv.Int())), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Float() * float64(rv.Int())), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Float() / float64(rv.Int())), nil
					}
				} else if bkl == reflect.Float64 && bkr == reflect.Uint64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Float() + float64(rv.Uint())), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Float() - float64(rv.Uint())), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Float() * float64(rv.Uint())), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Float() / float64(rv.Uint())), nil
					}
				} else if bkl == reflect.Float64 && bkr == reflect.Float64 {
					switch exp.MathOperator {
					case MathOperatorPlus:
						return reflect.ValueOf(lv.Float() + rv.Float()), nil
					case MathOperatorMinus:
						return reflect.ValueOf(lv.Float() - rv.Float()), nil
					case MathOperatorMul:
						return reflect.ValueOf(lv.Float() * rv.Float()), nil
					case MathOperatorDiv:
						return reflect.ValueOf(lv.Float() / rv.Float()), nil
					}
				}
			} else {
				return reflect.ValueOf(nil), fmt.Errorf("math operation can only be applied to numerical data, int, uit or float")
			}
		}
	}
	return reflect.ValueOf(nil), nil
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
