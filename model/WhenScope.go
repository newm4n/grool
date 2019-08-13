package model

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/pkg"
	"reflect"
)

type WhenScope struct {
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (when *WhenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	when.knowledgeContext = knowledgeContext
	when.ruleCtx = ruleCtx
	when.dataCtx = dataCtx

	if when.Expression != nil {
		when.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (when *WhenScope) AcceptExpression(expression *Expression) error {
	if when.Expression != nil {
		return fmt.Errorf("expression were set twice in when scope")
	}
	when.Expression = expression
	return nil
}

func (when *WhenScope) ExecuteWhen() (bool, error) {
	val, err := when.Expression.Evaluate()
	if err != nil {
		return false, errors.Trace(err)
	} else {
		if pkg.GetBaseKind(val) != reflect.Bool {
			return false, errors.Errorf("unexpected when result... its not boolean")
		} else {
			return val.Bool(), nil
		}
	}
}
