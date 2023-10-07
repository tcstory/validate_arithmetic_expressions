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

func toString(tokenList []Token) string {
	var text = ""

	for i := 0; i < len(tokenList); i++ {
		text = text + tokenList[i].value + ""
	}

	return fmt.Sprintf("%v", text)
}

type State struct {
	internal []string
}

func (self *State) push(ctx string) {
	self.internal = append(self.internal, ctx)
}

func (self *State) pop() string {
	if len(self.internal) >= 1 {
		var first = self.internal[0]

		self.internal = self.internal[1:]
		return first
	} else {
		return ""
	}
}

func (self *State) head() string {
	if len(self.internal) >= 1 {
		return self.internal[0]
	} else {
		return ""
	}
}

func (self *State) len() int {
	return len(self.internal)
}

type Result struct {
	value string
}

func run(tokenList []Token, state *State) (*Token, error) {
	// fmt.Printf("each %+v, state=%v\n\n", toString(tokenList), state)

	if len(tokenList) == 0 {
		return &Token{label: "operand", value: "_"}, nil
	} else if len(tokenList) == 1 {
		if tokenList[0].label == "operand" && state.head() != "leftParenthesis" {
			return &Token{
				label: "operand",
				value: "_",
			}, nil
		} else {
			return nil, errors.New("invalid expression")
		}
	} else {
		tokenList, operand1 := next(tokenList)

		if operand1.label == "leftParenthesis" {
			state.push("leftParenthesis")
			return run(tokenList, state)
		} else if operand1.label == "operand" {
			tokenList, operand2 := next(tokenList)
			tokenList, operand3 := next(tokenList)

			if operand2.label == "operator" && operand3.label == "operand" {
				return run(
					prepend(tokenList, &Token{label: "operand", value: "_"}),
					state,
				)
			} else if operand2.label == "operator" && operand3.label == "leftParenthesis" {
				state.push("leftParenthesis")
				var token, err = run(tokenList, state)
				if err != nil {
					return nil, err
				} else {
					return run([]Token{*operand1, *operand2, *token}, state)
				}
			} else if operand2.label == "rightParenthesis" && state.head() == "leftParenthesis" {
				if operand3 != nil {
					tokenList = prepend(tokenList, operand3)
				}

				state.pop()
				return run(
					prepend(
						tokenList,
						&Token{label: "operand", value: "_"},
					),
					state,
				)
			} else {
				return nil, errors.New("invalid expression")
			}
		} else {
			return nil, errors.New("invalid expression")
		}
	}
}

func Validate(input string) (*Result, error) {
	var tokenList = parse(input)

	var _, err = run(tokenList, &State{internal: []string{}})
	if err != nil {
		return nil, err
	} else {
		return &Result{value: "success"}, nil
	}
}
