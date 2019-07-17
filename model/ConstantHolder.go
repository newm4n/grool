package model

type ConstantHolder interface {
	AcceptConstant(cons *Constant) error
}
