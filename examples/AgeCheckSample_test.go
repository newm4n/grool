package examples

import (
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

	rule3 = `
rule AgeNameCheck "test"  salience 10{
    when
      Pogo.Compare(User.Name, "Calo")  
    then
      User.Name = "Success";
      Log(User.Name);
      Retract("AgeNameCheck");
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
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	err = dataContext.Add("Pogo", &MyPoGo{})
	if err != nil {
		t.Fatal(err)
	}

	//初始化规则引擎
	knowledgeBase := model.NewKnowledgeBase()
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase)

	err = ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(rule2)))
	if err != nil {
		t.Log(err)
	} else {
		eng1 := &engine.Grool{MaxCycle: 1}
		err := eng1.Execute(dataContext, knowledgeBase)
		if err != nil {
			t.Logf("Got error %v", err)
		} else {
			t.Log(user)
		}
	}
}

func TestMyPoGo_Compare(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	err = dataContext.Add("Pogo", &MyPoGo{})
	if err != nil {
		t.Fatal(err)
	}

	//初始化规则引擎
	knowledgeBase := model.NewKnowledgeBase()
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase)

	err = ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(rule3)))
	if err != nil {
		t.Log(err)
	} else {
		eng1 := &engine.Grool{MaxCycle: 100}
		err := eng1.Execute(dataContext, knowledgeBase)
		if err != nil {
			t.Logf("Got error %v", err)
		} else {
			t.Log(user)
		}
		if user.Name != "Success" {
			t.Logf("User should have changed name by rule to Success, but %s", user.Name)
			t.FailNow()
		}
	}
}
