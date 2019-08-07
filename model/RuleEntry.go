package model

import "github.com/newm4n/grool/context"

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

func (ins *RuleEntry) AcceptDecimal(val int64) error {
	ins.Salience = val
	return nil
}

func (ins *RuleEntry) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx

	if ins.WhenScope != nil {
		ins.WhenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ins.ThenScope != nil {
		ins.ThenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

func (re *RuleEntry) CanExecute() (bool, error) {
	bol, err := re.WhenScope.ExecuteWhen()
	if err != nil {
		return false, err
	}
	return bol, nil
}

func (re *RuleEntry) Execute() error {
	return re.ThenScope.Execute()
}
