package model

import (
	"github.com/newm4n/grool/context"
	"reflect"
)

type Constant struct {
	ConstantValue    reflect.Value
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (cons *Constant) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	cons.knowledgeContext = knowledgeContext
	cons.ruleCtx = ruleCtx
	cons.dataCtx = dataCtx
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (cons *Constant) Evaluate() (reflect.Value, error) {
	return cons.ConstantValue, nil
}

func (cons *Constant) AcceptDecimal(val int64) error {
	cons.ConstantValue = reflect.ValueOf(val)
	return nil
}
