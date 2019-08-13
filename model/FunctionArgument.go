package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"reflect"
)

type FunctionArgument struct {
	Arguments        []*ArgumentHolder
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// EvaluateArguments the object graph against underlined context or execute evaluation in the sub graph.
func (funcArg *FunctionArgument) EvaluateArguments() ([]reflect.Value, error) {
	if funcArg.Arguments == nil || len(funcArg.Arguments) == 0 {
		return make([]reflect.Value, 0), nil
	}
	retVal := make([]reflect.Value, len(funcArg.Arguments))
	for i, v := range funcArg.Arguments {
		rv, err := v.Evaluate()
		if err != nil {
			return retVal, errors.Trace(err)
		} else {
			retVal[i] = rv
		}
	}
	return retVal, nil
}

func (funcArg *FunctionArgument) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	funcArg.knowledgeContext = knowledgeContext
	funcArg.ruleCtx = ruleCtx
	funcArg.dataCtx = dataCtx

	if funcArg.Arguments != nil {
		for _, val := range funcArg.Arguments {
			val.Initialize(knowledgeContext, ruleCtx, dataCtx)
		}
	}
}

func (funcArg *FunctionArgument) AcceptExpression(expression *Expression) error {
	holder := &ArgumentHolder{
		Expression: expression,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

func (funcArg *FunctionArgument) AcceptFunctionCall(funcCall *FunctionCall) error {
	holder := &ArgumentHolder{
		FunctionCall: funcCall,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

func (funcArg *FunctionArgument) AcceptMethodCall(methodCall *MethodCall) error {
	holder := &ArgumentHolder{
		MethodCall: methodCall,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

func (funcArg *FunctionArgument) AcceptVariable(name string) error {
	holder := &ArgumentHolder{
		Variable: name,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

func (funcArg *FunctionArgument) AcceptConstant(cons *Constant) error {
	holder := &ArgumentHolder{
		Constant: cons,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}
