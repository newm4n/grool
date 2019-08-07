package model

type DecimalHolder interface {
	AcceptDecimal(val int64) error
}
