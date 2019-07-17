package model

type FunctionCallHolder interface {
	AcceptFunctionCall(funcCall *FunctionCall) error
}
