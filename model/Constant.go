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

func (ins *Constant) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ins *Constant) Evaluate() (reflect.Value, error) {
	return ins.ConstantValue, nil
}

func (ins *Constant) AcceptDecimal(val int64) error {
	ins.ConstantValue = reflect.ValueOf(val)
	return nil
}
