package model

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Assignment struct {
	Variable         string
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (assign *Assignment) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	assign.knowledgeContext = knowledgeContext
	assign.ruleCtx = ruleCtx
	assign.dataCtx = dataCtx

	if assign.Expression != nil {
		assign.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (assign *Assignment) AcceptExpression(expression *Expression) error {
	if assign.Expression != nil {
		return fmt.Errorf("expression were set twice in assignment")
	}
	assign.Expression = expression
	return nil
}

func (assign *Assignment) AcceptVariable(name string) error {
	if assign.Variable == "" {
		assign.Variable = name
		return nil
	} else {
		return errors.Errorf("variable already defined")
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (assign *Assignment) Evaluate() (reflect.Value, error) {
	v, err := assign.Expression.Evaluate()
	if err != nil {
		log.Errorf("Evaluate Got error %v", err)
		return reflect.ValueOf(nil), errors.Trace(err)
	}
	err = assign.dataCtx.SetValue(assign.Variable, v)
	if err != nil {
		log.Errorf("SetValue Got error %v", err)
		return reflect.ValueOf(nil), errors.Trace(err)
	}
	return reflect.ValueOf(nil), nil
}
