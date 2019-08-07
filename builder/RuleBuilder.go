package builder

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/juju/errors"
	antlr2 "github.com/newm4n/grool/antlr"
	"github.com/newm4n/grool/antlr/parser"
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
)

func NewRuleBuilder(KnowledgeBase *model.KnowledgeBase) *RuleBuilder {
	return &RuleBuilder{
		KnowledgeBase: KnowledgeBase,
	}
}

type RuleBuilder struct {
	KnowledgeBase *model.KnowledgeBase
}

func (builder *RuleBuilder) BuildRuleFromResources(resource []pkg.Resource) error {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(v)
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (builder *RuleBuilder) BuildRuleFromResource(resource pkg.Resource) error {
	data, err := resource.Load()
	if err != nil {
		return errors.Trace(err)
	}
	sdata := string(data)

	is := antlr.NewInputStream(sdata)
	lexer := parser.NewgroolLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := antlr2.NewGroolParserListener()

	parser := parser.NewgroolParser(stream)
	parser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Root())

	if len(listener.ParseErrors) > 0 {
		return errors.Errorf("error were found before builder bailing out. %d errors. 1st error : %v", len(listener.ParseErrors), listener.ParseErrors[0])
	} else {
		if len(listener.RuleEntries) > 0 {
			for k, v := range listener.RuleEntries {
				builder.KnowledgeBase.RuleEntries[k] = v
			}
		}
	}
	return nil
}
