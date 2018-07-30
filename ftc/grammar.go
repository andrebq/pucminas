package main

import (
	"errors"
)

type (
	grammarFn func(input tokens) (tokens, grammarFn, error)
)

func gramatica(input tokens) bool {
	return machine(input, programa)
}

func machine(input tokens, next grammarFn) bool {
	var err error
	for len(input) > 0 && err == nil && next != nil {
		input, next, err = next(input)
	}
	if err != nil {
		println("error: ", err.Error())
	}
	println("input", len(input), "err", err, "next", next)
	return len(input) == 0 && err == nil && next == nil
}

func matchHead(input tokens, kinds ...string) (tokens, tokens, bool) {
	oldInput := input
	var ret tokens
	for _, k := range kinds {
		if len(input) == 0 {
			break
		}
		if input[0].kind == k {
			ret = append(ret, input[0])
			input = input[1:]
		}
	}
	if len(ret) != len(kinds) {
		return ret, oldInput, false
	}
	return ret, input, true
}

func matchTail(input tokens, kinds ...string) (tokens, tokens, bool) {
	if len(input) < len(kinds) {
		return nil, input, false
	}
	tail := input[len(kinds)-1:]
	tail, _, ok := matchHead(tail)
	if !ok {
		return nil, input, false
	}
	return tail, input[:len(input)-len(kinds)], true
}

func programa(input tokens) (tokens, grammarFn, error) {
	tail := input
	_, tail, ok := matchHead(input, "inicio", "identificador", ";")
	if !ok {
		return input, nil, errors.New("missing inicio identificador ;")
	}
	_, tail, ok = matchTail(tail, "fim")
	if !ok {
		return input, nil, errors.New("missing fim")
	}
	return tail, declaracaoEInstrucao, nil
}

func declaracaoEInstrucao(input tokens) (tokens, grammarFn, error) {
	tail := input
	tail, _, err := declaracoes(tail)
	if err != nil {
		return input, nil, err
	}
	tail, _, err = instrucoes(tail)
	return tail, nil, err
}

func declaracoes(input tokens) (tokens, grammarFn, error) {
	println("declaracoes", input.String())
	tail := input
	tail, _, err := declaracao(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok := matchHead(tail, ";")
	if ok {
		return declaracoes(tail)
	}
	return tail, nil, nil
}

func declaracao(input tokens) (tokens, grammarFn, error) {
	println("declaracao", input.String())
	tail := input
	tail, _, err := identificadores(tail)
	if err != nil {
		return nil, nil, err
	}
	_, tail, ok := matchHead(tail, ":")
	if !ok {
		return nil, nil, errors.New("missing ':'")
	}
	tail, _, err = tipo(tail)

	return tail, nil, err
}

func identificadores(input tokens) (tokens, grammarFn, error) {
	tail := input
	tail, _, err := identificadorFn(tail)
	if err != nil {
		return nil, nil, err
	}
	_, tail, ok := matchHead(tail, ",")
	if ok {
		return identificadores(tail)
	}
	return tail, nil, nil
}

func identificadorFn(input tokens) (tokens, grammarFn, error) {
	tail := input
	_, tail, ok := matchHead(tail, "identificador")
	if !ok {
		return input, nil, errors.New("missing identificador")
	}
	return tail, nil, nil
}

func tipo(input tokens) (tokens, grammarFn, error) {
	tail := input
	for _, v := range []string{"inteiro", "real", "logico", "caracter"} {
		_, rest, ok := matchHead(tail, v)
		if ok {
			return rest, nil, nil
		}
	}
	return input, nil, errors.New("missing inteiro or real or logico or caracter")
}

func instrucoes(input tokens) (tokens, grammarFn, error) {
	tail := input
	tail, _, err := instrucao(tail)
	if err != nil {
		return tail, nil, err
	}
	_, tail, ok := matchHead(tail, ";")
	if ok {
		return instrucoes(tail)
	}
	return tail, nil, nil
}

func instrucao(input tokens) (tokens, grammarFn, error) {
	println("instrucao", input.String())
	opts := []grammarFn{
		atribuicaoFn,
		instrucaoSe,
		instrucaoEnquanto,
		instrucaoLeitura,
		instrucaoEscrita}
	tail := input
	for _, fn := range opts {
		rest, _, err := fn(tail)
		if err == nil {
			return rest, nil, nil
		}
	}
	return tail, nil, errors.New("missing a valid instrucao")
}

func instrucaoSe(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "se")
	if !ok {
		return input, nil, errors.New("missing 'se'")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, "entao", "inicio")
	if !ok {
		println("tail: ", tail.String())
		return input, nil, errors.New("missing 'entao inicio'")
	}
	tail, _, err = instrucoes(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, "fim")
	if !ok {
		return input, nil, errors.New("missing 'fim'")
	}
	_, tail, ok = matchHead(tail, "senao", "inicio")
	if ok {
		tail, _, err = instrucoes(tail)
		if err != nil {
			return input, nil, err
		}

		_, tail, ok = matchHead(tail, "fim")
		if !ok {
			return input, nil, errors.New("missing 'fim' from 'senao inicio'")
		}
	}
	return tail, nil, nil
}
func instrucaoLeitura(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "leia", "(")
	if !ok {
		return input, nil, errors.New("missing 'leia ('")
	}

	tail, _, err := identificadores(tail)
	if err != nil {
		return input, nil, err
	}

	_, tail, ok = matchHead(tail, ")")
	if !ok {
		return input, nil, errors.New("missing '('")
	}
	return tail, nil, nil
}

func instrucaoEscrita(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "escreva", "(")
	if !ok {
		return input, nil, errors.New("missing 'escreva ('")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, ")")
	if !ok {
		return input, nil, errors.New("missing ')'")
	}
	return tail, nil, nil
}

func instrucaoEnquanto(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "enquanto")
	if !ok {
		return input, nil, errors.New("missing 'enquanto'")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, "faca", "inicio")
	if !ok {
		return input, nil, errors.New("missing 'faca inicio'")
	}
	tail, _, err = instrucoes(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, "fim")
	if !ok {
		return input, nil, errors.New("missing 'fim'")
	}
	return tail, nil, nil
}

func atribuicaoFn(input tokens) (tokens, grammarFn, error) {
	println("atribuicao", input.String())
	tail := input
	tail, _, err := identificadorFn(tail)
	if err != nil {
		return tail, nil, err
	}
	_, tail, ok := matchHead(tail, ":=")
	if !ok {
		return tail, nil, errors.New("missing :=")
	}
	return expressao(tail)
}

func expressao(input tokens) (tokens, grammarFn, error) {
	opts := []grammarFn{
		expressaoConstante,
		expressaoIdentificador,
		expressaoParentesis}
	tail := input
	for _, fn := range opts {
		rest, _, err := fn(tail)
		if err == nil {
			tail = rest
			break
		}
	}

	// constante or identificador
	// maybe we need to look for the next expression and return it.
	if len(tail) == 0 {
		// no need to lookup further since there's nothing there anymore
		return tail, nil, nil
	}
	switch tail[0].kind {
	case "relop", "addop", "mulop", ",":
		opts = []grammarFn{
			expressaoLista,
			expressaoRel,
			expressaoAdd,
			expressaoMul,
		}
		for _, fn := range opts {
			rest, _, err := fn(tail)
			if err == nil {
				return rest, nil, nil
			}
		}
		return tail, nil, errors.New("missing a valid expressao")
	}
	return tail, nil, nil
}

func expressaoLista(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, ",")
	if !ok {
		return input, nil, errors.New("missing ','")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	return tail, nil, nil
}

func expressaoRel(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "relop")
	if !ok {
		return input, nil, errors.New("missing 'relop'")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	return tail, nil, nil
}
func expressaoAdd(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "addop")
	if !ok {
		return input, nil, errors.New("missing 'addop'")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	return tail, nil, nil
}
func expressaoMul(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "mulop")
	if !ok {
		return input, nil, errors.New("missing 'mulop'")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	return tail, nil, nil
}

func expressaoConstante(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "constante")
	if !ok {
		return nil, nil, errors.New("missing constante")
	}
	return tail, nil, nil
}

func expressaoParentesis(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "(")
	if !ok {
		return input, nil, errors.New("missing '('")
	}
	tail, _, err := expressao(tail)
	if err != nil {
		return input, nil, err
	}
	_, tail, ok = matchHead(tail, ")")
	if !ok {
		return input, nil, errors.New("missing ')'")
	}
	return tail, nil, nil
}
func expressaoIdentificador(input tokens) (tokens, grammarFn, error) {
	_, tail, ok := matchHead(input, "identificador")
	if !ok {
		return input, nil, errors.New("missing identificador")
	}
	return tail, nil, nil
}
