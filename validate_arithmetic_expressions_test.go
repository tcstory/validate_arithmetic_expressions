package validate_arithmetic_expressions

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	if len(parse("")) != 0 {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1"), []Token{
		{label: "operand", value: "1"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1+2"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "+"},
		{label: "operand", value: "2"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1-2"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "-"},
		{label: "operand", value: "2"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1*2"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "operand", value: "2"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1/2"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "/"},
		{label: "operand", value: "2"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1*(2+3)"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "leftParenthesis", value: "("},
		{label: "operand", value: "2"},
		{label: "operator", value: "+"},
		{label: "operand", value: "3"},
		{label: "rightParenthesis", value: ")"},
	}) {
		t.Fatalf("error found")
	}

	if !reflect.DeepEqual(parse("1*(2*(3+4))"), []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "leftParenthesis", value: "("},
		{label: "operand", value: "2"},
		{label: "operator", value: "*"},
		{label: "leftParenthesis", value: "("},
		{label: "operand", value: "3"},
		{label: "operator", value: "+"},
		{label: "operand", value: "4"},
		{label: "rightParenthesis", value: ")"},
		{label: "rightParenthesis", value: ")"},
	}) {
		t.Fatalf("error found")
	}
}

func TestNext(t *testing.T) {
	var tokenList = []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "operand", value: "2"},
	}
	var token *Token

	tokenList, token = next(tokenList)
	if *token != (Token{label: "operand", value: "1"}) {
		t.Fatalf("error found")
	}

	tokenList, token = next(tokenList)
	if *token != (Token{label: "operator", value: "*"}) {
		t.Fatalf("error found")
	}

	tokenList, token = next(tokenList)
	if *token != (Token{label: "operand", value: "2"}) {
		t.Fatalf("error found")
	}

	tokenList, token = next(tokenList)
	if token == nil && len(tokenList) == 0 {
		//
	} else {
		t.Fatalf("error found")
	}
}

func TestPrepend(t *testing.T) {
	var tokenList = []Token{
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "operand", value: "2"},
	}

	tokenList = prepend(tokenList, &Token{label: "operand", value: "0"})
	tokenList = prepend(tokenList, &Token{label: "operand", value: "-1"})

	if !reflect.DeepEqual(tokenList, []Token{
		{label: "operand", value: "-1"},
		{label: "operand", value: "0"},
		{label: "operand", value: "1"},
		{label: "operator", value: "*"},
		{label: "operand", value: "2"},
	}) {
		t.Fatalf("error found")
	}
}

func TestValidate(t *testing.T) {
	var _, err = Validate("1+2")

	if err != nil {
		t.Fatal(err)
	}

	_, err = Validate("1+2/3")

	if err != nil {
		t.Fatal(err)
	}

	_, err = Validate("")

	if err != nil {
		t.Fatal(err)
	}

}
