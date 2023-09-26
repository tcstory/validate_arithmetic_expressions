package validate_arithmetic_expressions

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
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
