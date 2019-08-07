package builder

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/juju/errors"
	antlr2 "github.com/newm4n/grool/antlr"
	"github.com/newm4n/grool/antlr/parser"
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
	log "github.com/sirupsen/logrus"
)

func NewRuleBuilder(KnowledgeBase *model.KnowledgeBase) *RuleBuilder {
	return &RuleBuilder{
		KnowledgeBase: KnowledgeBase,
	}
}

type RuleBuilder struct {
	KnowledgeBase *model.KnowledgeBase
}

func (builder *RuleBuilder) MustBuildRuleFromResources(resource []pkg.Resource) {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(v)
		if err != nil {
			panic(err)
		}
	}
}

func (builder *RuleBuilder) MustBuildRuleFromResource(resource pkg.Resource) {
	if err := builder.BuildRuleFromResource(resource); err != nil {
		panic(err)
	}
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

	listener := antlr2.NewGroolParserListener(builder.KnowledgeBase)

	psr := parser.NewgroolParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	if len(listener.ParseErrors) > 0 {
		log.Errorf("Loading rule resource : %s failed. Got %d errors. 1st error : %v", resource.String(), len(listener.ParseErrors), listener.ParseErrors[0])
		return errors.Errorf("error were found before builder bailing out. %d errors. 1st error : %v", len(listener.ParseErrors), listener.ParseErrors[0])
	} else {
		log.Debugf("Loading rule resource : %s success", resource.String())
		return nil
	}
}
