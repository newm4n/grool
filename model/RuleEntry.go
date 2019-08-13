package model

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
)

type RuleEntry struct {
	Salience         int64
	RuleName         string
	RuleDescription  string
	WhenScope        *WhenScope
	ThenScope        *ThenScope
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (entry *RuleEntry) AcceptDecimal(val int64) error {
	entry.Salience = val
	return nil
}

func (entry *RuleEntry) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	entry.knowledgeContext = knowledgeContext
	entry.ruleCtx = ruleCtx
	entry.dataCtx = dataCtx

	if entry.WhenScope != nil {
		entry.WhenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if entry.ThenScope != nil {
		entry.ThenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (entry *RuleEntry) CanExecute() (bool, error) {
	bol, err := entry.WhenScope.ExecuteWhen()
	if err != nil {
		return false, errors.Trace(err)
	}
	return bol, nil
}

func (entry *RuleEntry) Execute() error {
	return entry.ThenScope.Execute()
}
