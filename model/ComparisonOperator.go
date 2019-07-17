package model

var (
	ComparisonOperatorGT  = ComparisonOperator(">")
	ComparisonOperatorLT  = ComparisonOperator("<")
	ComparisonOperatorGTE = ComparisonOperator(">=")
	ComparisonOperatorLTE = ComparisonOperator("<=")
	ComparisonOperatorEQ  = ComparisonOperator("==")
	ComparisonOperatorNEQ = ComparisonOperator("!=")
)

type ComparisonOperator string
