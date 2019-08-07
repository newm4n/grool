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

func (ins *FunctionArgument) EvaluateArguments() ([]reflect.Value, error) {
	retVal := make([]reflect.Value, len(ins.Arguments))
	for i, v := range ins.Arguments {
		rv, err := v.Evaluate()
		if err != nil {
			return retVal, errors.Trace(err)
		} else {
			retVal[i] = rv
		}
	}
	return retVal, nil
}

func (ins *FunctionArgument) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.Arguments != nil {
		for _, val := range ins.Arguments {
			val.Initialize(knowledgeContext, ruleCtx, dataCtx)
		}
	}
}

func (arg *FunctionArgument) AcceptExpression(expression *Expression) error {
	holder := &ArgumentHolder{
		Expression: expression,
	}
	arg.Arguments = append(arg.Arguments, holder)
	return nil
}

func (arg *FunctionArgument) AcceptFunctionCall(funcCall *FunctionCall) error {
	holder := &ArgumentHolder{
		FunctionCall: funcCall,
	}
	arg.Arguments = append(arg.Arguments, holder)
	return nil
}

func (arg *FunctionArgument) AcceptVariable(name string) error {
	holder := &ArgumentHolder{
		Variable: name,
	}
	arg.Arguments = append(arg.Arguments, holder)
	return nil
}

func (arg *FunctionArgument) AcceptConstant(cons *Constant) error {
	holder := &ArgumentHolder{
		Constant: cons,
	}
	arg.Arguments = append(arg.Arguments, holder)
	return nil
}
