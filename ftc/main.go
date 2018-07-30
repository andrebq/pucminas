package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	help       = flag.Bool("h", false, "Exibe ajuda")
	tokensFlag = flag.String("tokens", "tokens.txt", "Arquivo de tokens")
	inputFile  = flag.String("input", "-", "input file (read from stdin by default)")
)

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(1)
	}
	input := readInput()

	ret, err := lex(bytes.NewBuffer(input))
	writeTokens(ret, err)
	if err != nil {
		fmt.Fprintf(os.Stdout, "NOK\r\n")
		os.Exit(1)
	}

	if !gramatica(ret) {
		fmt.Fprintf(os.Stdout, "NOK\r\n")
	} else {
		fmt.Fprintf(os.Stdout, "OK\r\n")
	}
}

func readInput() []byte {

	var file io.Reader
	if *inputFile == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(*inputFile)
		if err != nil {
			log.Fatalf("Erro ao abrir arquivo de entrada %v", err)
		}
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao efetuar a leitura da entrada: %v", err)
	}
	return buf
}

func writeTokens(ret tokens, err error) {
	file, err := os.Create(*tokensFlag)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de token para escrita: %v", err)
	}
	defer file.Close()

	for _, t := range ret {
		fmt.Fprintf(file, "<'%v' '%v' @ %v>\r\n", t.kind, t.value, t.pos)
	}

	if err != nil {
		fmt.Fprintf(file, "erro durante lex %v", err)
	}
}
