package helang

const (
	TokenKindNumber = iota
	TokenKindOr
	TokenKindIdent
	TokenKindLP
	TokenKindRP
	TokenKindLC
	TokenKindRC
	TokenKindLS
	TokenKindRS
	TokenKindComma
	TokenKindSemicolon
	TokenKindSub
	TokenKindIncrement
	TokenKindAssign
	TokenKindLT
	TokenKindKeyword
	TokenKindU8
	TokenKindPrint
	TokenKindSprint
	TokenKindTest5g
	TokenKindCyberspaces
)

var CharTokenKinds = map[byte]int {
	'|': TokenKindOr,
	'(': TokenKindLP,
	')': TokenKindRP,
	'{': TokenKindLC,
	'}': TokenKindRC,
	'[': TokenKindLS,
	']': TokenKindRS,
	',': TokenKindComma,
	';': TokenKindSemicolon,
	'=': TokenKindAssign,
	'<': TokenKindLT,
	'-': TokenKindSub,
}

var KeywordKinds = map[string]int {
	"print": TokenKindPrint,
	"u8": TokenKindU8,
	"test5g": TokenKindTest5g,
	"cyberspaces": TokenKindCyberspaces,
	"sprint": TokenKindSprint,
}

type TokenStruct struct {
	Content string
	Kind    int
}

func NewToken(content string, kind int) *TokenStruct {
	return &TokenStruct{
		Content: content, Kind: kind,
	}
}

func (token *TokenStruct) Compare (other *TokenStruct) bool {
	return token.Content == other.Content && token.Kind == other.Kind
}
