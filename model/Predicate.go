package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/pkg"
	"reflect"
	"time"
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

func (pred *Predicate) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if pred.ExpressionAtomLeft == nil {
		pred.ExpressionAtomLeft = exprAtom
	} else if pred.ExpressionAtomRight == nil {
		pred.ExpressionAtomRight = exprAtom
	} else {
		return errors.Errorf("expression alredy set twice")
	}
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (pre *Predicate) Evaluate() (reflect.Value, error) {
	if pre.ExpressionAtomRight == nil {
		return pre.ExpressionAtomLeft.Evaluate()
	} else {
		lv, err := pre.ExpressionAtomLeft.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		rv, err := pre.ExpressionAtomRight.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		if lv.Kind() == rv.Kind() && (pre.ComparisonOperator == ComparisonOperatorEQ || pre.ComparisonOperator == ComparisonOperatorNEQ) {
			if pre.ComparisonOperator == ComparisonOperatorEQ {
				switch lv.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					return reflect.ValueOf(lv.Int() == rv.Int()), nil
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					return reflect.ValueOf(lv.Uint() == rv.Uint()), nil
				case reflect.Float64, reflect.Float32:
					return reflect.ValueOf(lv.Float() == rv.Float()), nil
				case reflect.String:
					return reflect.ValueOf(lv.String() == rv.String()), nil
				case reflect.Bool:
					return reflect.ValueOf(lv.Bool() == rv.Bool()), nil
				}
				if lv.String() == "time.Time" {
					tl := pkg.ValueToInterface(lv).(time.Time)
					tr := pkg.ValueToInterface(rv).(time.Time)
					return reflect.ValueOf(tl.Equal(tr)), nil
				}
			} else {
				switch lv.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					return reflect.ValueOf(lv.Int() != rv.Int()), nil
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					return reflect.ValueOf(lv.Uint() != rv.Uint()), nil
				case reflect.Float64, reflect.Float32:
					return reflect.ValueOf(lv.Float() != rv.Float()), nil
				case reflect.String:
					return reflect.ValueOf(lv.String() != rv.String()), nil
				case reflect.Bool:
					return reflect.ValueOf(lv.Bool() != rv.Bool()), nil
				}
				if lv.String() == "time.Time" {
					tl := pkg.ValueToInterface(lv).(time.Time)
					tr := pkg.ValueToInterface(rv).(time.Time)
					return reflect.ValueOf(!tl.Equal(tr)), nil
				}
			}
		} else if lv.Type().String() == "time.Time" && rv.Type().String() == "time.Time" {
			tl := pkg.ValueToInterface(lv).(time.Time)
			tr := pkg.ValueToInterface(rv).(time.Time)
			switch pre.ComparisonOperator {
			case ComparisonOperatorEQ:
				return reflect.ValueOf(tl.Equal(tr)), nil
			case ComparisonOperatorNEQ:
				return reflect.ValueOf(!tl.Equal(tr)), nil
			case ComparisonOperatorGT:
				return reflect.ValueOf(tl.After(tr)), nil
			case ComparisonOperatorGTE:
				return reflect.ValueOf(tl.After(tr) || tl.Equal(tr)), nil
			case ComparisonOperatorLT:
				return reflect.ValueOf(tl.Before(tr)), nil
			case ComparisonOperatorLTE:
				return reflect.ValueOf(tl.Before(tr) || tl.Equal(tr)), nil
			}
		} else {
			var lf, rf float64
			if pkg.GetBaseKind(lv) == reflect.Int64 {
				lf = float64(lv.Int())
			} else if pkg.GetBaseKind(lv) == reflect.Uint64 {
				lf = float64(lv.Uint())
			} else if pkg.GetBaseKind(lv) == reflect.Float64 {
				lf = lv.Float()
			} else {
				return reflect.ValueOf(nil), errors.Errorf("comparison operator can only between strings, time or numbers")
			}
			if pkg.GetBaseKind(rv) == reflect.Int64 {
				rf = float64(rv.Int())
			} else if pkg.GetBaseKind(lv) == reflect.Uint64 {
				rf = float64(rv.Uint())
			} else if pkg.GetBaseKind(lv) == reflect.Float64 {
				rf = rv.Float()
			} else {
				return reflect.ValueOf(nil), errors.Errorf("comparison operator can only between strings, time or numbers")
			}
			switch pre.ComparisonOperator {
			case ComparisonOperatorEQ:
				return reflect.ValueOf(lf == rf), nil
			case ComparisonOperatorNEQ:
				return reflect.ValueOf(lf != rf), nil
			case ComparisonOperatorGT:
				return reflect.ValueOf(lf > rf), nil
			case ComparisonOperatorGTE:
				return reflect.ValueOf(lf >= rf), nil
			case ComparisonOperatorLT:
				return reflect.ValueOf(lf < rf), nil
			case ComparisonOperatorLTE:
				return reflect.ValueOf(lf <= rf), nil
			}
		}
	}
	return reflect.ValueOf(nil), nil
}
