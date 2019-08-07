package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
	"reflect"
)

type Expression struct {
	LeftExpression   *Expression
	RightExpression  *Expression
	LogicalOperator  LogicalOperator
	Predicate        *Predicate
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *Expression) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.LeftExpression != nil {
		ins.LeftExpression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.RightExpression != nil {
		ins.RightExpression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.Predicate != nil {
		ins.Predicate.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (expr *Expression) AcceptExpression(expression *Expression) error {
	if expr.LeftExpression != nil {
		if expr.RightExpression != nil {
			return fmt.Errorf("expression alredy set twice")
		}
		expr.RightExpression = expression
	}
	expr.LeftExpression = expression
	return nil
}

func (expr *Expression) Evaluate() (reflect.Value, error) {
	if expr.Predicate != nil {
		return expr.Predicate.Evaluate()
	} else {
		lv, err := expr.LeftExpression.Evaluate()
		if err != nil {
			return lv, err
		}
		rv, err := expr.RightExpression.Evaluate()
		if err != nil {
			return rv, err
		}
		if rv.Kind() == lv.Kind() && rv.Kind() == reflect.Bool {
			if expr.LogicalOperator == LogicalOperatorAnd {
				return reflect.ValueOf(lv.Bool() && rv.Bool()), nil
			} else {
				return reflect.ValueOf(lv.Bool() || rv.Bool()), nil
			}
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("cannot apply logical for non boolean expression")
		}
	}
}
