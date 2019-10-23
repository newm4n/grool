// Generated from grool.g4 by ANTLR 4.7.

package parser // grool

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 223,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 7, 2, 54, 10, 2, 12, 2,
	14, 2, 57, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 64, 10, 3, 3, 3,
	5, 3, 67, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 88, 10,
	9, 13, 9, 14, 9, 89, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 101, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 115, 10, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 7, 12, 121, 10, 12, 12, 12, 14, 12, 124, 11, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 131, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 144, 10, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 150, 10, 14, 12, 14, 14, 14, 153, 11,
	14, 3, 15, 3, 15, 3, 15, 5, 15, 158, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 5, 16, 165, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 174, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17,
	182, 10, 17, 7, 17, 184, 10, 17, 12, 17, 14, 17, 187, 11, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 5, 22, 204, 10, 22, 3, 22, 5, 22, 207, 10, 22, 3,
	23, 5, 23, 210, 10, 23, 3, 23, 3, 23, 3, 24, 5, 24, 215, 10, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 2, 4, 22, 26, 27, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 2, 8, 3, 2, 33, 34, 3, 2, 7, 8, 3, 2, 14, 15, 3, 2, 16, 19,
	4, 2, 20, 20, 22, 26, 3, 2, 9, 10, 2, 230, 2, 55, 3, 2, 2, 2, 4, 60, 3,
	2, 2, 2, 6, 73, 3, 2, 2, 2, 8, 76, 3, 2, 2, 2, 10, 78, 3, 2, 2, 2, 12,
	80, 3, 2, 2, 2, 14, 83, 3, 2, 2, 2, 16, 87, 3, 2, 2, 2, 18, 100, 3, 2,
	2, 2, 20, 102, 3, 2, 2, 2, 22, 114, 3, 2, 2, 2, 24, 130, 3, 2, 2, 2, 26,
	143, 3, 2, 2, 2, 28, 154, 3, 2, 2, 2, 30, 161, 3, 2, 2, 2, 32, 173, 3,
	2, 2, 2, 34, 188, 3, 2, 2, 2, 36, 190, 3, 2, 2, 2, 38, 192, 3, 2, 2, 2,
	40, 194, 3, 2, 2, 2, 42, 206, 3, 2, 2, 2, 44, 209, 3, 2, 2, 2, 46, 214,
	3, 2, 2, 2, 48, 218, 3, 2, 2, 2, 50, 220, 3, 2, 2, 2, 52, 54, 5, 4, 3,
	2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56,
	3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 7, 2, 2, 3,
	59, 3, 3, 2, 2, 2, 60, 61, 7, 4, 2, 2, 61, 63, 5, 8, 5, 2, 62, 64, 5, 10,
	6, 2, 63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 67,
	5, 6, 4, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 69, 7, 28, 2, 2, 69, 70, 5, 12, 7, 2, 70, 71, 5, 14, 8, 2, 71, 72,
	7, 29, 2, 2, 72, 5, 3, 2, 2, 2, 73, 74, 7, 13, 2, 2, 74, 75, 5, 44, 23,
	2, 75, 7, 3, 2, 2, 2, 76, 77, 7, 14, 2, 2, 77, 9, 3, 2, 2, 2, 78, 79, 9,
	2, 2, 2, 79, 11, 3, 2, 2, 2, 80, 81, 7, 5, 2, 2, 81, 82, 5, 22, 12, 2,
	82, 13, 3, 2, 2, 2, 83, 84, 7, 6, 2, 2, 84, 85, 5, 16, 9, 2, 85, 15, 3,
	2, 2, 2, 86, 88, 5, 18, 10, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2,
	89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 17, 3, 2, 2, 2, 91, 92, 5,
	20, 11, 2, 92, 93, 7, 27, 2, 2, 93, 101, 3, 2, 2, 2, 94, 95, 5, 28, 15,
	2, 95, 96, 7, 27, 2, 2, 96, 101, 3, 2, 2, 2, 97, 98, 5, 30, 16, 2, 98,
	99, 7, 27, 2, 2, 99, 101, 3, 2, 2, 2, 100, 91, 3, 2, 2, 2, 100, 94, 3,
	2, 2, 2, 100, 97, 3, 2, 2, 2, 101, 19, 3, 2, 2, 2, 102, 103, 5, 36, 19,
	2, 103, 104, 7, 21, 2, 2, 104, 105, 5, 22, 12, 2, 105, 21, 3, 2, 2, 2,
	106, 107, 8, 12, 1, 2, 107, 108, 7, 30, 2, 2, 108, 109, 5, 22, 12, 2, 109,
	110, 5, 34, 18, 2, 110, 111, 5, 22, 12, 2, 111, 112, 7, 31, 2, 2, 112,
	115, 3, 2, 2, 2, 113, 115, 5, 24, 13, 2, 114, 106, 3, 2, 2, 2, 114, 113,
	3, 2, 2, 2, 115, 122, 3, 2, 2, 2, 116, 117, 12, 5, 2, 2, 117, 118, 5, 34,
	18, 2, 118, 119, 5, 22, 12, 6, 119, 121, 3, 2, 2, 2, 120, 116, 3, 2, 2,
	2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	23, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126, 5, 26, 14, 2, 126, 127,
	5, 40, 21, 2, 127, 128, 5, 26, 14, 2, 128, 131, 3, 2, 2, 2, 129, 131, 5,
	26, 14, 2, 130, 125, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2, 131, 25, 3, 2, 2,
	2, 132, 133, 8, 14, 1, 2, 133, 144, 5, 42, 22, 2, 134, 144, 5, 36, 19,
	2, 135, 136, 7, 30, 2, 2, 136, 137, 5, 26, 14, 2, 137, 138, 5, 38, 20,
	2, 138, 139, 5, 26, 14, 2, 139, 140, 7, 31, 2, 2, 140, 144, 3, 2, 2, 2,
	141, 144, 5, 30, 16, 2, 142, 144, 5, 28, 15, 2, 143, 132, 3, 2, 2, 2, 143,
	134, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 142,
	3, 2, 2, 2, 144, 151, 3, 2, 2, 2, 145, 146, 12, 6, 2, 2, 146, 147, 5, 38,
	20, 2, 147, 148, 5, 26, 14, 7, 148, 150, 3, 2, 2, 2, 149, 145, 3, 2, 2,
	2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152,
	27, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 7, 15, 2, 2, 155, 157,
	7, 30, 2, 2, 156, 158, 5, 32, 17, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3,
	2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 7, 31, 2, 2, 160, 29, 3, 2, 2,
	2, 161, 162, 7, 14, 2, 2, 162, 164, 7, 30, 2, 2, 163, 165, 5, 32, 17, 2,
	164, 163, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	167, 7, 31, 2, 2, 167, 31, 3, 2, 2, 2, 168, 174, 5, 42, 22, 2, 169, 174,
	5, 36, 19, 2, 170, 174, 5, 30, 16, 2, 171, 174, 5, 28, 15, 2, 172, 174,
	5, 22, 12, 2, 173, 168, 3, 2, 2, 2, 173, 169, 3, 2, 2, 2, 173, 170, 3,
	2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 185, 3, 2, 2,
	2, 175, 181, 7, 3, 2, 2, 176, 182, 5, 42, 22, 2, 177, 182, 5, 36, 19, 2,
	178, 182, 5, 30, 16, 2, 179, 182, 5, 28, 15, 2, 180, 182, 5, 22, 12, 2,
	181, 176, 3, 2, 2, 2, 181, 177, 3, 2, 2, 2, 181, 178, 3, 2, 2, 2, 181,
	179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 184, 3, 2, 2, 2, 183, 175,
	3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2,
	2, 2, 186, 33, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 189, 9, 3, 2, 2,
	189, 35, 3, 2, 2, 2, 190, 191, 9, 4, 2, 2, 191, 37, 3, 2, 2, 2, 192, 193,
	9, 5, 2, 2, 193, 39, 3, 2, 2, 2, 194, 195, 9, 6, 2, 2, 195, 41, 3, 2, 2,
	2, 196, 207, 5, 48, 25, 2, 197, 207, 5, 44, 23, 2, 198, 199, 7, 17, 2,
	2, 199, 207, 5, 44, 23, 2, 200, 207, 5, 50, 26, 2, 201, 207, 5, 46, 24,
	2, 202, 204, 7, 12, 2, 2, 203, 202, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 207, 7, 11, 2, 2, 206, 196, 3, 2, 2, 2, 206, 197,
	3, 2, 2, 2, 206, 198, 3, 2, 2, 2, 206, 200, 3, 2, 2, 2, 206, 201, 3, 2,
	2, 2, 206, 203, 3, 2, 2, 2, 207, 43, 3, 2, 2, 2, 208, 210, 7, 17, 2, 2,
	209, 208, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211,
	212, 7, 35, 2, 2, 212, 45, 3, 2, 2, 2, 213, 215, 7, 17, 2, 2, 214, 213,
	3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 7, 36,
	2, 2, 217, 47, 3, 2, 2, 2, 218, 219, 9, 2, 2, 2, 219, 49, 3, 2, 2, 2, 220,
	221, 9, 7, 2, 2, 221, 51, 3, 2, 2, 2, 21, 55, 63, 66, 89, 100, 114, 122,
	130, 143, 151, 157, 164, 173, 181, 185, 203, 206, 209, 214,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "", "", "", "'&&'", "'||'", "", "", "", "", "", "", "", "'+'",
	"'-'", "'/'", "'*'", "'=='", "'='", "'>'", "'<'", "'>='", "'<='", "'!='",
	"';'", "'{'", "'}'", "'('", "')'", "'.'",
}
var symbolicNames = []string{
	"", "", "RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL",
	"NOT", "SALIENCE", "SIMPLENAME", "DOTTEDNAME", "PLUS", "MINUS", "DIV",
	"MUL", "EQUALS", "ASSIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "SEMICOLON",
	"LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING",
	"SQUOTA_STRING", "DECIMAL_LITERAL", "REAL_LITERAL", "SPACE", "COMMENT",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"root", "ruleEntry", "salience", "ruleName", "ruleDescription", "whenScope",
	"thenScope", "assignExpressions", "assignExpression", "assignment", "expression",
	"predicate", "expressionAtom", "methodCall", "functionCall", "functionArgs",
	"logicalOperator", "variable", "mathOperator", "comparisonOperator", "constant",
	"decimalLiteral", "realLiteral", "stringLiteral", "booleanLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type groolParser struct {
	*antlr.BaseParser
}

func NewgroolParser(input antlr.TokenStream) *groolParser {
	this := new(groolParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "grool.g4"

	return this
}

// groolParser tokens.
const (
	groolParserEOF             = antlr.TokenEOF
	groolParserT__0            = 1
	groolParserRULE            = 2
	groolParserWHEN            = 3
	groolParserTHEN            = 4
	groolParserAND             = 5
	groolParserOR              = 6
	groolParserTRUE            = 7
	groolParserFALSE           = 8
	groolParserNULL_LITERAL    = 9
	groolParserNOT             = 10
	groolParserSALIENCE        = 11
	groolParserSIMPLENAME      = 12
	groolParserDOTTEDNAME      = 13
	groolParserPLUS            = 14
	groolParserMINUS           = 15
	groolParserDIV             = 16
	groolParserMUL             = 17
	groolParserEQUALS          = 18
	groolParserASSIGN          = 19
	groolParserGT              = 20
	groolParserLT              = 21
	groolParserGTE             = 22
	groolParserLTE             = 23
	groolParserNOTEQUALS       = 24
	groolParserSEMICOLON       = 25
	groolParserLR_BRACE        = 26
	groolParserRR_BRACE        = 27
	groolParserLR_BRACKET      = 28
	groolParserRR_BRACKET      = 29
	groolParserDOT             = 30
	groolParserDQUOTA_STRING   = 31
	groolParserSQUOTA_STRING   = 32
	groolParserDECIMAL_LITERAL = 33
	groolParserREAL_LITERAL    = 34
	groolParserSPACE           = 35
	groolParserCOMMENT         = 36
	groolParserLINE_COMMENT    = 37
)

// groolParser rules.
const (
	groolParserRULE_root               = 0
	groolParserRULE_ruleEntry          = 1
	groolParserRULE_salience           = 2
	groolParserRULE_ruleName           = 3
	groolParserRULE_ruleDescription    = 4
	groolParserRULE_whenScope          = 5
	groolParserRULE_thenScope          = 6
	groolParserRULE_assignExpressions  = 7
	groolParserRULE_assignExpression   = 8
	groolParserRULE_assignment         = 9
	groolParserRULE_expression         = 10
	groolParserRULE_predicate          = 11
	groolParserRULE_expressionAtom     = 12
	groolParserRULE_methodCall         = 13
	groolParserRULE_functionCall       = 14
	groolParserRULE_functionArgs       = 15
	groolParserRULE_logicalOperator    = 16
	groolParserRULE_variable           = 17
	groolParserRULE_mathOperator       = 18
	groolParserRULE_comparisonOperator = 19
	groolParserRULE_constant           = 20
	groolParserRULE_decimalLiteral     = 21
	groolParserRULE_realLiteral        = 22
	groolParserRULE_stringLiteral      = 23
	groolParserRULE_booleanLiteral     = 24
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(groolParserEOF, 0)
}

func (s *RootContext) AllRuleEntry() []IRuleEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem())
	var tst = make([]IRuleEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntryContext)
		}
	}

	return tst
}

func (s *RootContext) RuleEntry(i int) IRuleEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleEntryContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *groolParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, groolParserRULE_root)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == groolParserRULE {
		{
			p.SetState(50)
			p.RuleEntry()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(groolParserEOF)
	}

	return localctx
}

// IRuleEntryContext is an interface to support dynamic dispatch.
type IRuleEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntryContext differentiates from other interfaces.
	IsRuleEntryContext()
}

type RuleEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntryContext() *RuleEntryContext {
	var p = new(RuleEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_ruleEntry
	return p
}

func (*RuleEntryContext) IsRuleEntryContext() {}

func NewRuleEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntryContext {
	var p = new(RuleEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_ruleEntry

	return p
}

func (s *RuleEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntryContext) RULE() antlr.TerminalNode {
	return s.GetToken(groolParserRULE, 0)
}

func (s *RuleEntryContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntryContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(groolParserLR_BRACE, 0)
}

func (s *RuleEntryContext) WhenScope() IWhenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhenScopeContext)
}

func (s *RuleEntryContext) ThenScope() IThenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThenScopeContext)
}

func (s *RuleEntryContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(groolParserRR_BRACE, 0)
}

func (s *RuleEntryContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntryContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterRuleEntry(s)
	}
}

func (s *RuleEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitRuleEntry(s)
	}
}

func (p *groolParser) RuleEntry() (localctx IRuleEntryContext) {
	localctx = NewRuleEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, groolParserRULE_ruleEntry)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(groolParserRULE)
	}
	{
		p.SetState(59)
		p.RuleName()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == groolParserDQUOTA_STRING || _la == groolParserSQUOTA_STRING {
		{
			p.SetState(60)
			p.RuleDescription()
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == groolParserSALIENCE {
		{
			p.SetState(63)
			p.Salience()
		}

	}
	{
		p.SetState(66)
		p.Match(groolParserLR_BRACE)
	}
	{
		p.SetState(67)
		p.WhenScope()
	}
	{
		p.SetState(68)
		p.ThenScope()
	}
	{
		p.SetState(69)
		p.Match(groolParserRR_BRACE)
	}

	return localctx
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(groolParserSALIENCE, 0)
}

func (s *SalienceContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitSalience(s)
	}
}

func (p *groolParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, groolParserRULE_salience)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(groolParserSALIENCE)
	}
	{
		p.SetState(72)
		p.DecimalLiteral()
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(groolParserSIMPLENAME, 0)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (p *groolParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, groolParserRULE_ruleName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(groolParserSIMPLENAME)
	}

	return localctx
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(groolParserDQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(groolParserSQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (p *groolParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, groolParserRULE_ruleDescription)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	_la = p.GetTokenStream().LA(1)

	if !(_la == groolParserDQUOTA_STRING || _la == groolParserSQUOTA_STRING) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IWhenScopeContext is an interface to support dynamic dispatch.
type IWhenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenScopeContext differentiates from other interfaces.
	IsWhenScopeContext()
}

type WhenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenScopeContext() *WhenScopeContext {
	var p = new(WhenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_whenScope
	return p
}

func (*WhenScopeContext) IsWhenScopeContext() {}

func NewWhenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenScopeContext {
	var p = new(WhenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_whenScope

	return p
}

func (s *WhenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenScopeContext) WHEN() antlr.TerminalNode {
	return s.GetToken(groolParserWHEN, 0)
}

func (s *WhenScopeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterWhenScope(s)
	}
}

func (s *WhenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitWhenScope(s)
	}
}

func (p *groolParser) WhenScope() (localctx IWhenScopeContext) {
	localctx = NewWhenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, groolParserRULE_whenScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(groolParserWHEN)
	}
	{
		p.SetState(79)
		p.expression(0)
	}

	return localctx
}

// IThenScopeContext is an interface to support dynamic dispatch.
type IThenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenScopeContext differentiates from other interfaces.
	IsThenScopeContext()
}

type ThenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenScopeContext() *ThenScopeContext {
	var p = new(ThenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_thenScope
	return p
}

func (*ThenScopeContext) IsThenScopeContext() {}

func NewThenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenScopeContext {
	var p = new(ThenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_thenScope

	return p
}

func (s *ThenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenScopeContext) THEN() antlr.TerminalNode {
	return s.GetToken(groolParserTHEN, 0)
}

func (s *ThenScopeContext) AssignExpressions() IAssignExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignExpressionsContext)
}

func (s *ThenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterThenScope(s)
	}
}

func (s *ThenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitThenScope(s)
	}
}

func (p *groolParser) ThenScope() (localctx IThenScopeContext) {
	localctx = NewThenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, groolParserRULE_thenScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(groolParserTHEN)
	}
	{
		p.SetState(82)
		p.AssignExpressions()
	}

	return localctx
}

// IAssignExpressionsContext is an interface to support dynamic dispatch.
type IAssignExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignExpressionsContext differentiates from other interfaces.
	IsAssignExpressionsContext()
}

type AssignExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignExpressionsContext() *AssignExpressionsContext {
	var p = new(AssignExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_assignExpressions
	return p
}

func (*AssignExpressionsContext) IsAssignExpressionsContext() {}

func NewAssignExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExpressionsContext {
	var p = new(AssignExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_assignExpressions

	return p
}

func (s *AssignExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExpressionsContext) AllAssignExpression() []IAssignExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignExpressionContext)(nil)).Elem())
	var tst = make([]IAssignExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignExpressionContext)
		}
	}

	return tst
}

func (s *AssignExpressionsContext) AssignExpression(i int) IAssignExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignExpressionContext)
}

func (s *AssignExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterAssignExpressions(s)
	}
}

func (s *AssignExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitAssignExpressions(s)
	}
}

func (p *groolParser) AssignExpressions() (localctx IAssignExpressionsContext) {
	localctx = NewAssignExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, groolParserRULE_assignExpressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == groolParserSIMPLENAME || _la == groolParserDOTTEDNAME {
		{
			p.SetState(84)
			p.AssignExpression()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignExpressionContext is an interface to support dynamic dispatch.
type IAssignExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignExpressionContext differentiates from other interfaces.
	IsAssignExpressionContext()
}

type AssignExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignExpressionContext() *AssignExpressionContext {
	var p = new(AssignExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_assignExpression
	return p
}

func (*AssignExpressionContext) IsAssignExpressionContext() {}

func NewAssignExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExpressionContext {
	var p = new(AssignExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_assignExpression

	return p
}

func (s *AssignExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExpressionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignExpressionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(groolParserSEMICOLON, 0)
}

func (s *AssignExpressionContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *AssignExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterAssignExpression(s)
	}
}

func (s *AssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitAssignExpression(s)
	}
}

func (p *groolParser) AssignExpression() (localctx IAssignExpressionContext) {
	localctx = NewAssignExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, groolParserRULE_assignExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Assignment()
		}
		{
			p.SetState(90)
			p.Match(groolParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.MethodCall()
		}
		{
			p.SetState(93)
			p.Match(groolParserSEMICOLON)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.FunctionCall()
		}
		{
			p.SetState(96)
			p.Match(groolParserSEMICOLON)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(groolParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *groolParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, groolParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Variable()
	}
	{
		p.SetState(101)
		p.Match(groolParserASSIGN)
	}
	{
		p.SetState(102)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(groolParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(groolParserRR_BRACKET, 0)
}

func (s *ExpressionContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *groolParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *groolParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, groolParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(105)
			p.Match(groolParserLR_BRACKET)
		}
		{
			p.SetState(106)
			p.expression(0)
		}
		{
			p.SetState(107)
			p.LogicalOperator()
		}
		{
			p.SetState(108)
			p.expression(0)
		}
		{
			p.SetState(109)
			p.Match(groolParserRR_BRACKET)
		}

	case 2:
		{
			p.SetState(111)
			p.Predicate()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, groolParserRULE_expression)
			p.SetState(114)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(115)
				p.LogicalOperator()
			}
			{
				p.SetState(116)
				p.expression(4)
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *PredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *PredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *groolParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, groolParserRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.expressionAtom(0)
		}
		{
			p.SetState(124)
			p.ComparisonOperator()
		}
		{
			p.SetState(125)
			p.expressionAtom(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.expressionAtom(0)
		}

	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionAtomContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionAtomContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionAtomContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionAtomContext)

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionAtomContext
	right  IExpressionAtomContext
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *ExpressionAtomContext) GetRight() IExpressionAtomContext { return s.right }

func (s *ExpressionAtomContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *ExpressionAtomContext) SetRight(v IExpressionAtomContext) { s.right = v }

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(groolParserLR_BRACKET, 0)
}

func (s *ExpressionAtomContext) MathOperator() IMathOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathOperatorContext)
}

func (s *ExpressionAtomContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(groolParserRR_BRACKET, 0)
}

func (s *ExpressionAtomContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *ExpressionAtomContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (p *groolParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	return p.expressionAtom(0)
}

func (p *groolParser) expressionAtom(_p int) (localctx IExpressionAtomContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, groolParserRULE_expressionAtom, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(131)
			p.Constant()
		}

	case 2:
		{
			p.SetState(132)
			p.Variable()
		}

	case 3:
		{
			p.SetState(133)
			p.Match(groolParserLR_BRACKET)
		}
		{
			p.SetState(134)

			var _x = p.expressionAtom(0)

			localctx.(*ExpressionAtomContext).left = _x
		}
		{
			p.SetState(135)
			p.MathOperator()
		}
		{
			p.SetState(136)

			var _x = p.expressionAtom(0)

			localctx.(*ExpressionAtomContext).right = _x
		}
		{
			p.SetState(137)
			p.Match(groolParserRR_BRACKET)
		}

	case 4:
		{
			p.SetState(139)
			p.FunctionCall()
		}

	case 5:
		{
			p.SetState(140)
			p.MethodCall()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
			localctx.(*ExpressionAtomContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, groolParserRULE_expressionAtom)
			p.SetState(143)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(144)
				p.MathOperator()
			}
			{
				p.SetState(145)

				var _x = p.expressionAtom(5)

				localctx.(*ExpressionAtomContext).right = _x
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(groolParserDOTTEDNAME, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (p *groolParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, groolParserRULE_methodCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(groolParserDOTTEDNAME)
	}
	{
		p.SetState(153)
		p.Match(groolParserLR_BRACKET)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-7)&-(0x1f+1)) == 0 && ((1<<uint((_la-7)))&((1<<(groolParserTRUE-7))|(1<<(groolParserFALSE-7))|(1<<(groolParserNULL_LITERAL-7))|(1<<(groolParserNOT-7))|(1<<(groolParserSIMPLENAME-7))|(1<<(groolParserDOTTEDNAME-7))|(1<<(groolParserMINUS-7))|(1<<(groolParserLR_BRACKET-7))|(1<<(groolParserDQUOTA_STRING-7))|(1<<(groolParserSQUOTA_STRING-7))|(1<<(groolParserDECIMAL_LITERAL-7))|(1<<(groolParserREAL_LITERAL-7)))) != 0 {
		{
			p.SetState(154)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(157)
		p.Match(groolParserRR_BRACKET)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(groolParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *groolParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, groolParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(groolParserSIMPLENAME)
	}
	{
		p.SetState(160)
		p.Match(groolParserLR_BRACKET)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-7)&-(0x1f+1)) == 0 && ((1<<uint((_la-7)))&((1<<(groolParserTRUE-7))|(1<<(groolParserFALSE-7))|(1<<(groolParserNULL_LITERAL-7))|(1<<(groolParserNOT-7))|(1<<(groolParserSIMPLENAME-7))|(1<<(groolParserDOTTEDNAME-7))|(1<<(groolParserMINUS-7))|(1<<(groolParserLR_BRACKET-7))|(1<<(groolParserDQUOTA_STRING-7))|(1<<(groolParserSQUOTA_STRING-7))|(1<<(groolParserDECIMAL_LITERAL-7))|(1<<(groolParserREAL_LITERAL-7)))) != 0 {
		{
			p.SetState(161)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(164)
		p.Match(groolParserRR_BRACKET)
	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgsContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionArgsContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionArgsContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (p *groolParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, groolParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(166)
			p.Constant()
		}

	case 2:
		{
			p.SetState(167)
			p.Variable()
		}

	case 3:
		{
			p.SetState(168)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(169)
			p.MethodCall()
		}

	case 5:
		{
			p.SetState(170)
			p.expression(0)
		}

	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == groolParserT__0 {
		{
			p.SetState(173)
			p.Match(groolParserT__0)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(174)
				p.Constant()
			}

		case 2:
			{
				p.SetState(175)
				p.Variable()
			}

		case 3:
			{
				p.SetState(176)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(177)
				p.MethodCall()
			}

		case 5:
			{
				p.SetState(178)
				p.expression(0)
			}

		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(groolParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(groolParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (p *groolParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, groolParserRULE_logicalOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	_la = p.GetTokenStream().LA(1)

	if !(_la == groolParserAND || _la == groolParserOR) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(groolParserSIMPLENAME, 0)
}

func (s *VariableContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(groolParserDOTTEDNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *groolParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, groolParserRULE_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	_la = p.GetTokenStream().LA(1)

	if !(_la == groolParserSIMPLENAME || _la == groolParserDOTTEDNAME) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IMathOperatorContext is an interface to support dynamic dispatch.
type IMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathOperatorContext differentiates from other interfaces.
	IsMathOperatorContext()
}

type MathOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathOperatorContext() *MathOperatorContext {
	var p = new(MathOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_mathOperator
	return p
}

func (*MathOperatorContext) IsMathOperatorContext() {}

func NewMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathOperatorContext {
	var p = new(MathOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_mathOperator

	return p
}

func (s *MathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(groolParserMUL, 0)
}

func (s *MathOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(groolParserDIV, 0)
}

func (s *MathOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(groolParserPLUS, 0)
}

func (s *MathOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(groolParserMINUS, 0)
}

func (s *MathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterMathOperator(s)
	}
}

func (s *MathOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitMathOperator(s)
	}
}

func (p *groolParser) MathOperator() (localctx IMathOperatorContext) {
	localctx = NewMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, groolParserRULE_mathOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<groolParserPLUS)|(1<<groolParserMINUS)|(1<<groolParserDIV)|(1<<groolParserMUL))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(groolParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(groolParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(groolParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(groolParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(groolParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(groolParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *groolParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, groolParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(192)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<groolParserEQUALS)|(1<<groolParserGT)|(1<<groolParserLT)|(1<<groolParserGTE)|(1<<groolParserLTE)|(1<<groolParserNOTEQUALS))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(groolParserNULL_LITERAL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(groolParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *groolParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, groolParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.DecimalLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.Match(groolParserMINUS)
		}
		{
			p.SetState(197)
			p.DecimalLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.BooleanLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(199)
			p.RealLiteral()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == groolParserNOT {
			{
				p.SetState(200)
				p.Match(groolParserNOT)
			}

		}
		{
			p.SetState(203)
			p.Match(groolParserNULL_LITERAL)
		}

	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(groolParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(groolParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (p *groolParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, groolParserRULE_decimalLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == groolParserMINUS {
		{
			p.SetState(206)
			p.Match(groolParserMINUS)
		}

	}
	{
		p.SetState(209)
		p.Match(groolParserDECIMAL_LITERAL)
	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(groolParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(groolParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (p *groolParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, groolParserRULE_realLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == groolParserMINUS {
		{
			p.SetState(211)
			p.Match(groolParserMINUS)
		}

	}
	{
		p.SetState(214)
		p.Match(groolParserREAL_LITERAL)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(groolParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(groolParserSQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *groolParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, groolParserRULE_stringLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	_la = p.GetTokenStream().LA(1)

	if !(_la == groolParserDQUOTA_STRING || _la == groolParserSQUOTA_STRING) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = groolParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = groolParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(groolParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(groolParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(groolListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *groolParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, groolParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	_la = p.GetTokenStream().LA(1)

	if !(_la == groolParserTRUE || _la == groolParserFALSE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

func (p *groolParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 12:
		var t *ExpressionAtomContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionAtomContext)
		}
		return p.ExpressionAtom_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *groolParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *groolParser) ExpressionAtom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
