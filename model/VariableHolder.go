package model

type VariableHolder interface {
	AcceptVariable(name string) error
}
