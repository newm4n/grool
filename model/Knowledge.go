package model

type KnowledgeBase struct {
	RuleEntries map[string]*RuleEntry
}

func NewKnowledgeBase() *KnowledgeBase {
	return &KnowledgeBase{
		RuleEntries: make(map[string]*RuleEntry),
	}
}

func (k *KnowledgeBase) Retract(ruleEntryName string) {
	if re, ok := k.RuleEntries[ruleEntryName]; ok {
		re.Retracted = true
	}
}

func (k *KnowledgeBase) Reset() {
	for _, v := range k.RuleEntries {
		v.Retracted = false
	}
}
