package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/newm4n/grool/antlr/parser"
	"io/ioutil"
	"testing"
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
			output := fmt.Sprintf("%s (%q)\n", lexer.SymbolicNames[nt.GetTokenType()], nt.GetText())
			fmt.Print(output)
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

		listener := &GroolParserListener{}

		parser := parser.NewgroolParser(stream)
		parser.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, parser.Root())

		for _, e := range listener.ParseErrors {
			t.Log(e)
			t.FailNow()
		}
	}

}
