package model

type ExpressionAtomHolder interface {
	AcceptExpressionAtom(exprAtom *ExpressionAtom) error
}
