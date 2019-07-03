package antlr

import (
	"fmt"
	"github.com/newm4n/grool/antlr/parser"
)

type GroolParserListener struct {
	parser.BasegroolListener
	ParseErrors []error
}

func (s *GroolParserListener) AddError(e error) {
	s.ParseErrors = append(s.ParseErrors, e)
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GroolParserListener) EnterRuleEntry(ctx *parser.RuleEntryContext) {
	fmt.Println("")
}
