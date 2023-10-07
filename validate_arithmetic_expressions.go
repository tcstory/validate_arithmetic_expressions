package validate_arithmetic_expressions

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
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

func next(tokenList []Token) ([]Token, *Token) {
	if len(tokenList) > 0 {
		var first = tokenList[0]
		return slices.Delete(tokenList, 0, 1), &first
	} else {
		return []Token{}, nil
	}
}

func prepend(tokenList []Token, token *Token) []Token {
	return append([]Token{*token}, tokenList...)
}

type Result struct {
	value string
}

func Validate(input string) (Result, error) {
	var tokenList = parse(input)
	var run func(tokenList []Token) (Result, error)

	run = func(tokenList []Token) (Result, error) {

		if len(tokenList) == 0 {
			return Result{
				value: "success",
			}, nil
		} else if len(tokenList) == 1 {
			if tokenList[0].label == "operand" {
				return Result{
					value: "success",
				}, nil
			} else {
				return Result{}, errors.New("wrong label type")
			}
		} else {
			tokenList, operand1 := next(tokenList)
			tokenList, operator := next(tokenList)
			tokenList, operand2 := next(tokenList)

			if operand1.label == "operand" && operator.label == "operator" && operand2.label == "operand" {
				return run(prepend(tokenList, &Token{label: "operand", value: ""}))
			} else {
				return Result{}, errors.New("invalid expression")
			}
		}
	}

	return run(tokenList)
}
