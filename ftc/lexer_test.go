package main

import (
	"bytes"
	"testing"
)

type (
	lexTestCase struct {
		input  string
		output tokens
	}

	checker func(tokens, tokens) bool
)

func checkIgnorePosition(a, b tokens) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, v := range a {
		if v.kind != b[idx].kind || v.value != b[idx].value {
			return false
		}
	}
	return true
}

func runTest(t *testing.T, c checker, tests ...lexTestCase) {
	for _, test := range tests {
		ret, err := lex(bytes.NewBufferString(test.input))
		if err != nil {
			t.Fatal(err)
		}
		if !c(ret, test.output) {
			t.Errorf("for [%v] should get [%v] got [%v]",
				test.input, test.output, ret)
		}
	}
}

func TestDigito(t *testing.T) {
	runTest(t, checkIgnorePosition, lexTestCase{
		input: "123 456 789",
		output: tokens{
			{value: "123", kind: "digito"},
			{value: "456", kind: "digito"},
			{value: "789", kind: "digito"},
		},
	})
}

func TestIdentificadorEKeyworkd(t *testing.T) {
	runTest(t, checkIgnorePosition, lexTestCase{
		input: "abc1 123 inicio fim senao",
		output: tokens{
			{value: "abc1", kind: "identificador"},
			{value: "123", kind: "digito"},
			{value: "inicio", kind: "inicio"},
			{value: "fim", kind: "fim"},
			{value: "senao", kind: "senao"},
		},
	})
}

func TestRelop(t *testing.T) {
	runTest(t, checkIgnorePosition, lexTestCase{
		input: "= << <= > >= <>",
		output: tokens{
			{value: "=", kind: "relop"},
			{value: "<", kind: "relop"},
			// syntatically it is invalid to have two << in
			// a sequence, but to the lexer that sequence is fine
			// if we had bitwise operations, that that could be
			// ambiguous
			{value: "<", kind: "relop"},
			{value: "<=", kind: "relop"},
			{value: ">", kind: "relop"},
			{value: ">=", kind: "relop"},
			{value: "<>", kind: "relop"},
		},
	})
}
