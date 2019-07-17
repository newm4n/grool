package model

type ExpressionHolder interface {
	AcceptExpression(expression *Expression) error
}
