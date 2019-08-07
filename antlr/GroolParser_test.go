package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/newm4n/grool/antlr/parser"
	"github.com/newm4n/grool/model"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestLexer(t *testing.T) {
	data, err := ioutil.ReadFile("./sample.grool")
	if err != nil {
		t.Fatal(err)
	} else {
		is := antlr.NewInputStream(string(data))

		// Create the Lexer
		lexer := parser.NewgroolLexer(is)
		//lexer := parser.NewLdifParserLexer(is)

		// Read all tokens
		for {
			nt := lexer.NextToken()
			if nt.GetTokenType() == antlr.TokenEOF {
				break
			}
		}
	}

}

func TestParser(t *testing.T) {
	data, err := ioutil.ReadFile("./sample.grool")
	if err != nil {
		t.Fatal(err)
	} else {
		sdata := string(data)

		is := antlr.NewInputStream(sdata)
		lexer := parser.NewgroolLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		listener := NewGroolParserListener(model.NewKnowledgeBase())

		psr := parser.NewgroolParser(stream)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

		for _, e := range listener.ParseErrors {
			t.Log(e)
			t.FailNow()
		}
	}

}

func TestTimeKind(t *testing.T) {
	n := time.Now()
	nt := reflect.TypeOf(n)
	fmt.Println(nt.String())
}
