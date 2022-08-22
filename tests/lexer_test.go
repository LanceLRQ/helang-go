package tests

import (
	assertProvider "github.com/stretchr/testify/assert"
	"helang-go/helang"
	"helang-go/helang/core"
	"testing"
)

func TestComments(t *testing.T) {
	assert := assertProvider.New(t)
	code := `
		u8 a = 1 | // Comment inline.
		2
		// Comment for single line.
	`
	lexer := helang.NewLex([]byte(code))
	tokens, err := lexer.Lex()
	if err != nil {
		t.Fatal(err)
	}
	expectTokens := helang.TokensList {
		core.NewToken("u8", core.TokenKindU8),
		core.NewToken("a", core.TokenKindIdent),
		core.NewToken("=", core.TokenKindAssign),
		core.NewToken("1", core.TokenKindNumber),
		core.NewToken("|", core.TokenKindOr),
		core.NewToken("2", core.TokenKindNumber),
	}
	assert.Equal(expectTokens, tokens)
}
