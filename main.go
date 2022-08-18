package main

import (
	"fmt"
	"helang-go/helang"
	"io/ioutil"
)

func main() {
	hFile, err := ioutil.ReadFile("./great.he")
	if err != nil {
		panic(err)
	}
	lexer := helang.NewLex(hFile)
	tokens, err := lexer.Lex()
	if err != nil {
		panic(err)
	}
	for _, token := range tokens {
		fmt.Printf("content: %s, kind: %d\n", token.Content, token.Kind)
	}
}
