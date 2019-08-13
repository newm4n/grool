package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

type MethodCall struct {
	MethodName       string
	MethodArguments  *FunctionArgument
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *MethodCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.MethodArguments != nil {
		ins.MethodArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (ins *MethodCall) AcceptFunctionArgument(funcArg *FunctionArgument) error {
	ins.MethodArguments = funcArg
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (exp *MethodCall) Evaluate() (reflect.Value, error) {
	var argumentValues []reflect.Value
	if exp.MethodArguments == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := exp.MethodArguments.EvaluateArguments()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		argumentValues = av
	}

	return exp.dataCtx.ExecMethod(exp.MethodName, argumentValues)
}
