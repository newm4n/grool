package examples

import (
	"fmt"
	"github.com/newm4n/grool/builder"
	"github.com/newm4n/grool/context"
	"github.com/newm4n/grool/engine"
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
	"testing"
)

const (
	rule2 = `
rule AgeNameCheck "test" {
    when
      Pogo.GetStringLength("9999") > 0 
    then
      Log(User.Name);
}
`
)

func TestMyPoGo_GetStringLength(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	dataContext.Add("User", user)
	dataContext.Add("Pogo", &MyPoGo{})

	//初始化规则引擎
	knowledgeBase := model.NewKnowledgeBase()
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase)

	err := ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(rule2)))
	if err != nil {
		fmt.Println(err)
	} else {
		eng1 := &engine.Grool{MaxCycle: 1}
		eng1.Execute(dataContext, knowledgeBase)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(user)
		}
	}
}
