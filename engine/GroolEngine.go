package engine

import (
	"github.com/juju/errors"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/model"
	"sort"
)

func NewGroolEngine() *Grool {
	return &Grool{}
}

type Grool struct {
}

func (g *Grool) Execute(dataCtx *context.DataContext, knowledge *model.KnowledgeBase) error {
	kctx := &context.KnowledgeContext{}
	rctx := &context.RuleContext{}
	for _, v := range knowledge.RuleEntries {
		v.Initialize(kctx, rctx, dataCtx)
	}

	for true {
		runnable := make([]*model.RuleEntry, 0)
		for _, v := range knowledge.RuleEntries {
			can, err := v.CanExecute()
			if err != nil {
				return errors.Trace(err)
			}
			if can {
				runnable = append(runnable, v)
			}
		}
		if len(runnable) > 0 {
			if len(runnable) > 1 {
				sort.SliceStable(runnable, func(i, j int) bool {
					return runnable[i].Salience > runnable[j].Salience
				})
			}
			err := runnable[0].Execute()
			if err != nil {
				return errors.Trace(err)
			}
		} else {
			break
		}
	}
	return nil
}
