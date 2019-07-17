package model

import (
	"github.com/newm4n/grool/context"
	"time"
)

type Constant struct {
	StringValue      string
	DecimalValue     int64
	BooleanValue     bool
	FloatValue       float64
	TimeValue        time.Time
	IsNull           bool
	DataType         DataType
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

func (ins *Constant) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ins.knowledgeContext = knowledgeContext
	ins.ruleCtx = ruleCtx
	ins.dataCtx = dataCtx
}
