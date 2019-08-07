package model

import "reflect"

type Evaluator interface {
	Evaluate() (reflect.Value, error)
}

type Executor interface {
	Execute() error
}
