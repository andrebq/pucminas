package main

import (
	"fmt"
	"unicode/utf8"
)

func initialFn(input *reader, e *emitter) stateFn {
	r := input.readRune()
	if utf8.RuneError == r {
		return nil
	}
	if isWhitespace(r) {
		return initialFn
	} else if isLetra(r) {
		e.push(r)
		return identificadorOuKeyword
	} else if isNumero(r) {
		e.push(r)
		return digito
	} else if isRelop(r) {
		input.pushback(r)
		return relop
	} else if isAddop(r) {
		input.pushback(r)
		return addop
	} else if isMulop(r) {
		input.pushback(r)
		return mulop
	} else if r == '(' {
		e.push(r)
		e.emit("(", input)
		return initialFn
	} else if r == ')' {
		e.push(r)
		e.emit(")", input)
		return initialFn
	} else if r == ';' {
		e.push(r)
		e.emit(";", input)
		return initialFn
	} else if r == ',' {
		e.push(r)
		e.emit(",", input)
		return initialFn
	} else if r == ':' {
		input.pushback(r)
		return atribuicao
	}
	println("rune: ", r, "str", fmt.Sprintf("%q", string(r)))
	return nil
}

func isAddop(r rune) bool {
	return r == '+' || r == '-'
}

func isMulop(r rune) bool {
	return r == '*' || r == '/'
}

func atribuicao(input *reader, e *emitter) stateFn {
	colon := input.readRune()
	if colon != ':' {
		return nil
	}
	equal := input.readRune()
	if isWhitespace(equal) {
		input.pushback(equal)
		e.push(colon)
		e.emit(":", input)
		return initialFn
	}
	e.push(colon)
	e.push(equal)
	e.emit(":=", input)
	return initialFn
}

func mulop(input *reader, e *emitter) stateFn {
	r := input.readRune()
	switch r {
	case '*':
		e.push(r)
		e.emit("mulop", input)
		return initialFn
	case '/':
		e.push(r)
		e.emit("mulop", input)
		return initialFn
	}
	return nil
}

func addop(input *reader, e *emitter) stateFn {
	r := input.readRune()
	switch r {
	case '+', '-':
		e.push(r)
	}
	d := input.readRune()
	if isNumero(d) {
		input.pushback(d)
		return digito
	} else {
		e.emit("addop", input)
		input.pushback(d)
		return initialFn
	}
	return nil
}

func relop(input *reader, e *emitter) stateFn {
	r := input.readRune()
	switch r {
	case '=':
		e.push(r)
		e.emit("relop", input)
		return initialFn
	case '>':
		e.push(r)
		return relopMaiorOuIgual
	case '<':
		e.push(r)
		return relopMenorOuIgualOuDiferente
	}
	input.pushback(r)
	return initialFn
}

func relopMaiorOuIgual(input *reader, e *emitter) stateFn {
	r := input.readRune()
	switch r {
	case '=':
		e.push(r)
		e.emit("relop", input)
	default:
		input.pushback(r)
		e.emit("relop", input)
	}
	return initialFn
}

func relopMenorOuIgualOuDiferente(input *reader, e *emitter) stateFn {
	r := input.readRune()
	switch r {
	case '=', '>':
		e.push(r)
		e.emit("relop", input)
	default:
		input.pushback(r)
		e.emit("relop", input)
	}
	return initialFn
}

func identificador(input *reader, e *emitter) stateFn {
	r := input.readRune()
	if isLetra(r) {
		e.push(r)
		return identificadorOuKeyword
	} else if isNumero(r) {
		e.push(r)
		return identificador
	}
	e.emit("identificador", input)
	input.pushback(r)
	return initialFn
}

func identificadorOuKeyword(input *reader, e *emitter) stateFn {
	r := input.readRune()
	if isLetra(r) {
		e.push(r)
		return identificadorOuKeyword
	} else if isNumero(r) {
		e.push(r)
		// since we have a number
		// there is no way for this to be a keyword anymore
		// return the correct state in this case
		return identificador
	}

	str := string(e.acc)
	switch str {
	case "inicio", "fim", "inteiro", "real", "logico", "caracter",
		"se", "entao", "senao", "leia", "escreva", "enquanto",
		"faca":
		e.emit(str, input)
	case "or":
		e.emit("addop", input)
	case "div", "mod", "and":
		e.emit("mulop", input)
	default:
		e.emit("identificador", input)
	}
	input.pushback(r)
	return initialFn
}

func digito(input *reader, e *emitter) stateFn {
	r := input.readRune()
	if isNumero(r) {
		e.push(r)
		return digito
	}
	e.emit("constante", input)
	// since we don't now how to handle
	// runes that aren't digitis
	// push it back to the reader
	// and let the initialState figure it out
	input.pushback(r)
	return initialFn
}
