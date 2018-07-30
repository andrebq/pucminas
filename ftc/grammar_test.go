package main

import (
	"bytes"
	"testing"
)

func TestProgramaValido(t *testing.T) {
	ret, err := lex(bytes.NewBufferString(`
		inicio teste;

		idade, saida : real; salario : inteiro;
		
		letra : caracter
		
		a := +4;
		b := a;
		
		se ( idade  >= -9) entao inicio
		 enquanto (b = +10) and (c <> -50)
		 faca inicio
		
		   leia(idade);
		   escreva(salario)
		 fim
		fim;

		se (10) entao inicio leia(idade) fim
		senao inicio escreva(idade) fim; 
		
		leia (saida)
		
		fim  		
		`))
	if err != nil {
		t.Fatal(err, ret)
	}

	if !gramatica(ret) {
		t.Fatal("programa deveria ser valido")
	}
}
