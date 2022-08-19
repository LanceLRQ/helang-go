package main

import (
	_ "embed"
	"helang-go/helang"
	"helang-go/helang/core"
	"log"
)

//go:embed great.he
var hFile string

func main() {

	lexer := helang.NewLex([]byte(hFile))
	tokens, err := lexer.Lex()
	if err != nil {
		log.Fatal(err)
	}

	parser := helang.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	env := map[string]*core.U8{}
	_, err = ast.Evaluate(env)
	if err != nil {
		log.Fatal(err)
	}
}
