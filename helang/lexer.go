package helang

import (
	"fmt"
	"regexp"
)

const (
	LexerStateWait = iota
	LexerStateIdent
	LexerStateNumber
	LexerStateIncrement
	LexerStateComment
)

type TokensList []*TokenStruct

type Lexer struct {
	content []byte
	contentLength int
	state int
	pos int
	cache string
	cur byte
	line int
	lineChar int
	tokens TokensList
}

func NewLex (content []byte) *Lexer {
	lexer := &Lexer {
		content: content,
		state: LexerStateWait,
		pos: -1,
		lineChar: 0,
		line: 1,
		cache: "",
		cur: 0,
	}
	return lexer
}

func (lexer *Lexer) Lex() (TokensList, error) {
	lexer.tokens = make(TokensList, 0, 10)
	lexer.pos = -1
	lexer.line = 1
	lexer.lineChar = 0
	lexer.nextChar()
	var err error
	for lexer.pos < len(lexer.content) {
		switch lexer.state {
		case LexerStateWait:
			err = lexer.wait()
			break
		case LexerStateIdent:
			lexer.ident()
			break
		case LexerStateNumber:
			lexer.number()
			break
		case LexerStateIncrement:
			err = lexer.increment()
			break
		case LexerStateComment:
			err = lexer.comment()
			break
		}
		if err != nil {
			return lexer.tokens, err
		}
	}
	return lexer.tokens, nil
}

func (lexer *Lexer) nextChar() {
	lexer.pos++
	if lexer.pos >= len(lexer.content) {
		return
	}
	lexer.cur = lexer.content[lexer.pos]
	if lexer.cur == '\n' {
		lexer.line++
		lexer.lineChar = 0
	} else {
		lexer.lineChar++
	}
}

func (lexer *Lexer) wait() error {
	// Anyway, clear the cache.
	lexer.cache = ""

	if matched, _ := regexp.Match("\\s", []byte{lexer.cur}); matched {
		//  Matched space, skipping.
		lexer.nextChar()
		return nil
	}

	if lexer.cur == '/' {
		lexer.state = LexerStateComment
		return nil
	}

	if matched, _ := regexp.Match("\\d", []byte{lexer.cur}); matched {
		// Matched number, changing state to NUMBER.
		lexer.state = LexerStateNumber
		return nil
	}

	if matched, _ := regexp.Match("[a-zA-Z_$]", []byte{lexer.cur}); matched {
		// Matched identifier, changing state to IDENT.
		lexer.state = LexerStateIdent
		return nil
	}

	if lexer.cur == '+' {
		// Matched increment operator, changing state to INCREMENT.
		lexer.state = LexerStateIncrement
		return nil
	}

	for char, kind := range CharTokenKinds {
		if lexer.cur == char {
			// Matched single char token, adding it to the list.
			lexer.tokens = append(lexer.tokens, NewToken(string(lexer.cur), kind))
			lexer.nextChar()
			return nil
		}
	}
	return fmt.Errorf("bad token '%s' (line %d:%d)", string(lexer.cur), lexer.line, lexer.lineChar)
}

func (lexer *Lexer) ident() {
	matched, _ := regexp.Match("[A-Za-z0-9_$]", []byte{lexer.cur})
	if lexer.cache != "" && !matched {
		// Current character is not identifier, changing state to WAIT.
		kind, ok := KeywordKinds[lexer.cache]
		if !ok {
			kind = TokenKindIdent
			lexer.tokens = append(lexer.tokens, NewToken(lexer.cache, kind))
			lexer.state = LexerStateWait
			return
		}
	}
	lexer.cache += string(lexer.cur)
	lexer.nextChar()
	return
}

func (lexer *Lexer) number() {
	// Not support for floats yet, as the King He hasn't written any floats.
	matched, _ := regexp.Match("\\d", []byte{lexer.cur})
	if !matched {
		// Current character is not number, changing state to WAIT.
		lexer.tokens = append(lexer.tokens, NewToken(lexer.cache, TokenKindNumber))
		lexer.state = LexerStateWait
		return
	}
	lexer.cache += string(lexer.cur)
	lexer.nextChar()
	return
}

func (lexer *Lexer) increment() error {
	// Not support for floats yet, as the King He hasn't written any floats.
	if lexer.cache == "+" && lexer.cur != '+' {
		return fmt.Errorf("only ++ operator is expected, as the King He has NOT written single + (line %d:%d)", lexer.line, lexer.lineChar)
	}
	if lexer.cache == "++" {
		// Enough + operator, changing state to WAIT.
		lexer.tokens = append(lexer.tokens, NewToken(lexer.cache, TokenKindIncrement))
		lexer.state = LexerStateWait
		return nil
	}
	lexer.cache += string(lexer.cur)
	lexer.nextChar()
	return nil
}


func (lexer *Lexer) comment() error {
	if lexer.cache == "/" && lexer.cur != '/' {
		return fmt.Errorf("bad token '%s' (line %d:%d)", string(lexer.cur), lexer.line, lexer.lineChar)
	}
	if lexer.cache == "//" {
		if lexer.cur == '\n' {
			lexer.state = LexerStateWait
		}
		lexer.nextChar()
		return nil
	}
	lexer.cache += string(lexer.cur)
	lexer.nextChar()
	return nil
}
