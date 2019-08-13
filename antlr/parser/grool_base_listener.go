// Generated from grool.g4 by ANTLR 4.7.

package parser // grool

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegroolListener is a complete listener for a parse tree produced by groolParser.
type BasegroolListener struct{}

var _ groolListener = &BasegroolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegroolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegroolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegroolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegroolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasegroolListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasegroolListener) ExitRoot(ctx *RootContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *BasegroolListener) EnterRuleEntry(ctx *RuleEntryContext) {}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *BasegroolListener) ExitRuleEntry(ctx *RuleEntryContext) {}

// EnterSalience is called when production salience is entered.
func (s *BasegroolListener) EnterSalience(ctx *SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *BasegroolListener) ExitSalience(ctx *SalienceContext) {}

// EnterRuleName is called when production ruleName is entered.
func (s *BasegroolListener) EnterRuleName(ctx *RuleNameContext) {}

// ExitRuleName is called when production ruleName is exited.
func (s *BasegroolListener) ExitRuleName(ctx *RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *BasegroolListener) EnterRuleDescription(ctx *RuleDescriptionContext) {}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *BasegroolListener) ExitRuleDescription(ctx *RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *BasegroolListener) EnterWhenScope(ctx *WhenScopeContext) {}

// ExitWhenScope is called when production whenScope is exited.
func (s *BasegroolListener) ExitWhenScope(ctx *WhenScopeContext) {}

// EnterThenScope is called when production thenScope is entered.
func (s *BasegroolListener) EnterThenScope(ctx *ThenScopeContext) {}

// ExitThenScope is called when production thenScope is exited.
func (s *BasegroolListener) ExitThenScope(ctx *ThenScopeContext) {}

// EnterAssignExpressions is called when production assignExpressions is entered.
func (s *BasegroolListener) EnterAssignExpressions(ctx *AssignExpressionsContext) {}

// ExitAssignExpressions is called when production assignExpressions is exited.
func (s *BasegroolListener) ExitAssignExpressions(ctx *AssignExpressionsContext) {}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *BasegroolListener) EnterAssignExpression(ctx *AssignExpressionContext) {}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *BasegroolListener) ExitAssignExpression(ctx *AssignExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasegroolListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasegroolListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasegroolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasegroolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasegroolListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasegroolListener) ExitPredicate(ctx *PredicateContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *BasegroolListener) EnterExpressionAtom(ctx *ExpressionAtomContext) {}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *BasegroolListener) ExitExpressionAtom(ctx *ExpressionAtomContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BasegroolListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BasegroolListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasegroolListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasegroolListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BasegroolListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BasegroolListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BasegroolListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BasegroolListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasegroolListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasegroolListener) ExitVariable(ctx *VariableContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BasegroolListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BasegroolListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BasegroolListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BasegroolListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasegroolListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasegroolListener) ExitConstant(ctx *ConstantContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BasegroolListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BasegroolListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *BasegroolListener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *BasegroolListener) ExitRealLiteral(ctx *RealLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasegroolListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasegroolListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BasegroolListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BasegroolListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
