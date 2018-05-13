package main

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

type (
	runes  []rune
	tokens []token

	token struct {
		value string
		kind  string
		pos   pos
	}

	emitter struct {
		acc    runes
		tokens tokens
	}

	reader struct {
		input io.Reader
		lines int
		col   int

		eof bool

		top rune
	}

	pos struct {
		line int
		col  int
	}

	stateFn func(*reader, *emitter) stateFn
)

func lex(input io.Reader) (tokens, error) {
	emit := emitter{
		tokens: nil,
	}

	reader := newReader(input)

	err := dontPanic(func() {
		var state stateFn
		state = initialFn
		for state != nil {
			state = state(reader, &emit)
		}
	})

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		return nil, err
	}

	if !reader.eof {
		return emit.tokens, errors.New("found an nil state before reaching EOF")
	}

	return emit.tokens, nil
}

func (p pos) back(cols int) pos {
	return pos{
		line: p.line,
		col:  p.col - cols,
	}
}

func (e *emitter) emit(kind string, p *reader) {
	e.tokens = append(e.tokens, token{
		kind:  kind,
		value: string(e.acc),
		pos:   p.pos().back(len(e.acc)),
	})
	e.acc = nil
}

func (e *emitter) push(r rune) {
	e.acc = append(e.acc, r)
}

func newReader(input io.Reader) *reader {
	return &reader{
		input: input,
		lines: 0,
		col:   0,
		top:   utf8.RuneError,
	}
}

func (r *reader) pushback(v rune) {
	if v == '\n' || v == '\r' {
		// ignore pushbacks for newlines
		// since they are irrelevant to our lexer
		return
	}
	r.top = v
	r.col--
}

func (r *reader) readRune() rune {
	if r.eof {
		panic(io.EOF)
	}

	if r.top != utf8.RuneError {
		var ret rune
		ret, r.top = r.top, utf8.RuneError
		r.col++
		return ret
	}
	var buf [1]byte
	_, err := r.input.Read(buf[:])
	if err == io.EOF {
		r.eof = true
		return utf8.RuneError
	}
	if err != nil {
		panic(err)
	}
	runeVal, _ := utf8.DecodeRune(buf[:])
	if runeVal == utf8.RuneError {
		panic(fmt.Errorf("invalid encoding use ASCII only"))
	}

	if runeVal == '\n' {
		r.lines++
		r.col = 0
	} else if runeVal != '\r' {
		r.col++
	} // do nothing on \r

	return runeVal
}

func (r *reader) pos() pos {
	return pos{line: r.lines, col: r.col}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r'
}

func isLetra(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumero(r rune) bool {
	return r >= '0' && r <= '9'
}

func isRelop(r rune) bool {
	return r == '=' || r == '>' || r == '<'
}

func dontPanic(fn func()) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = fmt.Errorf("recover: %v", rec)
			}
		}
	}()
	fn()
	return err
}
