package model

type KnowledgeBase struct {
	RuleEntries map[string]*RuleEntry
}

func NewKnowledgeBase() *KnowledgeBase {
	return &KnowledgeBase{
		RuleEntries: make(map[string]*RuleEntry),
	}
}
