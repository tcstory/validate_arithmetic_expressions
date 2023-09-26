package validate_arithmetic_expressions

import (
	"fmt"
	"regexp"
)

type Token struct {
	label string
	value string
}

func (self *Token) toString() string {
	return fmt.Sprintf(
		"{label: %q, token: %q}", self.label, self.value,
	)
}

func parse(input string) []Token {
	var reOfOperand = regexp.MustCompile(`\d+`)
	var reOfOperator = regexp.MustCompile(`[+\-*/]`)

	var list = []Token{}

	for _, ch := range input {
		var label string
		var value = string(ch)

		if reOfOperand.MatchString(value) {
			label = "operand"
		} else if reOfOperator.MatchString(value) {
			label = "operator"
		} else if value == "(" {
			label = "leftParenthesis"
		} else {
			label = "rightParenthesis"
		}

		list = append(list, Token{
			label: label,
			value: value,
		})
	}

	return list
}
