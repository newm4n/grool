package model

type FunctionArgumentHolder interface {
	AcceptFunctionArgument(funcArg *FunctionArgument) error
}
