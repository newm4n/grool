package model

import (
	"fmt"
	"github.com/newm4n/grool/context"
	"reflect"
)

// ArgumentHolder is a struct part of the rule object graph.
// It holds child graph such as Variable name, Constant data, Function, Expressions, etc.
type ArgumentHolder struct {
	Constant         *Constant
	Variable         string
	FunctionCall     *FunctionCall
	MethodCall       *MethodCall
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// Initialize this ArgumentHolder instance graph before rule execution start.
func (ins *ArgumentHolder) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.Constant != nil {
		ins.Constant.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.FunctionCall != nil {
		ins.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.MethodCall != nil {
		ins.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ins.Expression != nil {
		ins.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ins *ArgumentHolder) Evaluate() (reflect.Value, error) {
	if len(ins.Variable) > 0 {
		return ins.dataCtx.GetValue(ins.Variable)
	} else if ins.Constant != nil {
		return ins.Constant.Evaluate()
	} else if ins.FunctionCall != nil {
		return ins.FunctionCall.Evaluate()
	} else if ins.Expression != nil {
		return ins.Expression.Evaluate()
	} else {
		return reflect.ValueOf(nil), fmt.Errorf("argument holder stores no value")
	}
}
