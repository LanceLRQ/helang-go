package helang

import (
	assertProvider "github.com/stretchr/testify/assert"
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
	lexer := NewLex([]byte(code))
	tokens, err := lexer.Lex()
	if err != nil {
		t.Fatal(err)
	}
	expectTokens := TokensList{
		core.NewToken("u8", core.TokenKindU8),
		core.NewToken("a", core.TokenKindIdent),
		core.NewToken("=", core.TokenKindAssign),
		core.NewToken("1", core.TokenKindNumber),
		core.NewToken("|", core.TokenKindOr),
		core.NewToken("2", core.TokenKindNumber),
	}
	assert.Equal(expectTokens, tokens)
}
