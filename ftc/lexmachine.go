package main

func initialFn(input *reader, e *emitter) stateFn {
	r := input.readRune()
	if isWhitespace(r) {
		return initialFn
	} else if isLetra(r) {
		e.push(r)
		return identificadorOuKeyword
	} else if isNumero(r) {
		e.push(r)
		return digito
	} else if r == '+' {
		e.push(r)
		e.emit("plus", input)
	} else if r == '-' {
		e.push('-')
		e.emit("negative", input)
	} else if isRelop(r) {
		input.pushback(r)
		return relop
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
	e.emit("digito", input)
	// since we don't now how to handle
	// runes that aren't digitis
	// push it back to the reader
	// and let the initialState figure it out
	input.pushback(r)
	return initialFn
}
