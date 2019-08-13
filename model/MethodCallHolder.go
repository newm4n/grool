package model

type MethodCallHolder interface {
	AcceptMethodCall(methodCall *MethodCall) error
}
