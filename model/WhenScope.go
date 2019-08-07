package model

import (
	"fmt"
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

func (ins *WhenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.Expression != nil {
		ins.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ws *WhenScope) AcceptExpression(expression *Expression) error {
	if ws.Expression != nil {
		return fmt.Errorf("expression were set twice in when scope")
	}
	ws.Expression = expression
	return nil
}

func (ws *WhenScope) ExecuteWhen() (bool, error) {
	val, err := ws.Expression.Evaluate()
	if err != nil {
		return false, err
	} else {
		if pkg.GetBaseKind(val) != reflect.Bool {
			return false, fmt.Errorf("unexpected when result... its not boolean")
		} else {
			return val.Bool(), nil
		}
	}
}
