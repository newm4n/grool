package model

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

type FunctionCall struct {
	FunctionName      string
	FunctionArguments *FunctionArgument
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *context.RuleContext
	dataCtx           *context.DataContext
}

func (ins *FunctionCall) AcceptFunctionArgument(funcArg *FunctionArgument) error {
	ins.FunctionArguments = funcArg
	return nil
}

func (ins *FunctionCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.FunctionArguments != nil {
		ins.FunctionArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (exp *FunctionCall) Evaluate() (reflect.Value, error) {
	var argumentValues []reflect.Value
	if exp.FunctionArguments == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := exp.FunctionArguments.EvaluateArguments()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		argumentValues = av
	}

	fName := fmt.Sprintf("DEFUNC.%s", exp.FunctionName)
	return exp.dataCtx.ExecMethod(fName, argumentValues)
}
