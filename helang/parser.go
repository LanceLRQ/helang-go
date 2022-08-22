package helang

import (
	"errors"
	"fmt"
	"helang-go/helang/asts"
	"helang-go/helang/core"
	"strconv"
)

type ParserFuncType func(*Parser, ...interface{}) (asts.BasicAST, error)

type Parser struct {
	tokens TokensList
	pos int
}

func NewParser(tokens TokensList) *Parser {
	return &Parser{
		tokens: tokens,
		pos: 0,
	}
}

var rootParsers = []ParserFuncType {
	_rootParsePrint,
	_rootParseSprint,
	_rootParseU8Set,
	_rootParseVarDef,
	_rootParseVarDeclare,
	_rootParseVarAssign,
	_rootParseVarIncrement,
	_rootParseExpr,
	_rootParseTest5g,
	_rootParseSemicolon,
	_rootParseCyberspaces,
}

func (parser *Parser) Parse() (asts.BasicAST, error) {
	/*
	 root
	          : print
	          | sprint
	          | u8_set
	          | var_def
	          | var_declare
	          | var_assign
	          | var_increment
	          | expr
	          | test_5g
	          | semicolon
	          | cyberspaces
	          ;
	 */
	heAsts := make([]asts.BasicAST, 0, 1)
	for parser.pos < len(parser.tokens) {
		pyForElse := true
		for _, ep := range rootParsers {
			savedPos := parser.pos
			rel, err := ep(parser)
			if err == nil {
				heAsts = append(heAsts, rel)
				pyForElse = false
				break
			} else if errors.Is(err, core.BadStatementException) {
				parser.pos = savedPos
			}
		}
		if pyForElse {
			return nil, fmt.Errorf(
				"%w: failed to parse tokens started from %d, which is %s",
				core.BadStatementException,
				parser.pos,
				parser.tokens[parser.pos].String(),
			)
		}
	}

	// Return the AST itself if there is only one.
	if len(heAsts) != 1 {
		return asts.NewListAST(heAsts), nil
	}
	return heAsts[0], nil
}

func (parser *Parser) _expect(expectedKind int) (*core.TokenStruct, error) {
	if parser.pos >= len(parser.tokens) {
		return nil, fmt.Errorf("%w: no more tokens", core.BadStatementException)
	}
	token := parser.tokens[parser.pos]

	if token.Kind != expectedKind {
		return nil, fmt.Errorf("%w: expected %d at pos %d, got %d", core.BadStatementException, expectedKind, parser.pos, token.Kind)
	}

	parser.pos++

	return token, nil
}

// == ast parser

func _rootParseCyberspaces(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// cyberspaces: CYBERSPACES SEMICOLON;

	_, err := parser._expect(core.TokenKindCyberspaces)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewCyberspacesAST(), nil
}

func _rootParseSemicolon(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	_, err := parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewVoidAST(), nil
}

func _rootParseVarDef(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// var_def: U8 IDENT ASSIGN expr SEMICOLON;

	_, err := parser._expect(core.TokenKindU8)
	if err != nil { return nil, err }

	varIdent, err := parser._expect(core.TokenKindIdent)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindAssign)
	if err != nil { return nil, err }

	val, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewVarDefAST(varIdent.Content, val), nil
}

func _rootParseVarDeclare(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// var_declare: U8 IDENT SEMICOLON;
	_, err := parser._expect(core.TokenKindU8)
	if err != nil { return nil, err }

	varIdent, err := parser._expect(core.TokenKindIdent)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }
	return asts.NewVarDefAST(varIdent.Content, asts.NewVoidAST()), nil
}

func _rootParseVarAssign(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// var_assign: IDENT ASSIGN expr SEMICOLON;

	varIdent, err := parser._expect(core.TokenKindIdent)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindAssign)
	if err != nil { return nil, err }

	val, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewVarDefAST(varIdent.Content, val), nil
}

func _rootParsePrint(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// print: PRINT expr SEMICOLON;

	_, err := parser._expect(core.TokenKindPrint)
	if err != nil { return nil, err }

	expr, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewPrintAST(expr), nil
}

func _rootParseSprint(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// print: SPRINT expr SEMICOLON;

	_, err := parser._expect(core.TokenKindSprint)
	if err != nil { return nil, err }

	expr, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }
	return asts.NewSprintAST(expr), nil
}

func _rootParseVarIncrement(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// var_increment: IDENT INCREMENT SEMICOLON;

	ident, err := parser._expect(core.TokenKindIdent)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindIncrement)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewVarIncrementAST(ident.Content), nil
}

func _rootParseTest5g(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
		test_5g: TEST_5G SEMICOLON;
	*/
	_, err := parser._expect(core.TokenKindTest5g)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }
	return asts.NewTest5GAST(), nil
}

func _parseU8CommonParts(parser *Parser, args ...interface{}) ([]asts.BasicAST, error) {
	listExpr, err := _rootParseExpr(parser, true)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindLS)
	if err != nil { return nil, err }

	subscriptExpr, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindRS)
	if err != nil { return nil, err }

	return []asts.BasicAST {listExpr, subscriptExpr}, nil
}

func _rootParseU8Set(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
	    u8_set: expr LS expr RS ASSIGN expr SEMICOLON;
	*/
	expr, err := _parseU8CommonParts(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindAssign)
	if err != nil { return nil, err }

	val, err := _rootParseExpr(parser)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindSemicolon)
	if err != nil { return nil, err }

	return asts.NewU8SetAST(expr[0], expr[1], val), nil
}

func _rootParseExpr(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
	   expr
	      : empty_u8_expr
	      | or_u8_expr
	      | var_expr
	      | u8_get_expr
	      ;
	    :param skip_u8: whether it should skip u8.
	*/
	skipU8 := false
	if len(args) > 0 {
		skipU8 = args[0].(bool)
	}

	exprParsers := []ParserFuncType {
		_parseEmptyU8Expr,
		_parseOrU8Expr,
		_parseU8GetExpr,
		_parseVarExpr,
	}

	for i, ep := range exprParsers {
		if skipU8 && i == 2 { // ep == _parseU8GetExpr
			continue
		}
		savedPos := parser.pos
		rel, err := ep(parser)
		if err == nil {
			return rel, nil
		} else if errors.Is(err, core.BadStatementException) {
			parser.pos = savedPos
		}
	}

	return nil, fmt.Errorf("%w: cannot parse expressions", core.BadStatementException)
}

func _parseEmptyU8Expr(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	// empty_u8_expr: LS NUMBER RS;
	_, err := parser._expect(core.TokenKindLS)
	if err != nil { return nil, err }

	length, err := parser._expect(core.TokenKindNumber)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindRS)
	if err != nil { return nil, err }

	aLen, err := strconv.ParseInt(length.Content, 10, 32)
	if err != nil { return nil, err }

	return asts.NewEmptyU8InitAST(int(aLen)), nil
}

func _parseOrU8Expr(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
	    or_u8_expr
            : NUMBER
            | NUMBER OR or_u8_expr
            ;
	*/
	first, err := parser._expect(core.TokenKindNumber)
	if err != nil { return nil, err }

	firstVal, err := strconv.ParseInt(first.Content, 10, 32)
	if err != nil { return nil, err }

	_, err = parser._expect(core.TokenKindOr)
	if err != nil {
		if errors.Is(err, core.BadStatementException) {
			return asts.NewOrU8InitAST(int(firstVal), nil), nil
		}
		return nil, err
	}
	second, err := _parseOrU8Expr(parser)
	if err != nil { return nil, err }

	return asts.NewOrU8InitAST(int(firstVal), second.(*asts.OrU8InitAST)), nil
}

func _parseVarExpr(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
		var_expr: IDENT;
	*/
	ident, err := parser._expect(core.TokenKindIdent)
	if err != nil { return nil, err }

	return asts.NewVarExprAST(ident.Content), nil
}

func _parseU8GetExpr(parser *Parser, args ...interface{}) (asts.BasicAST, error) {
	/*
		u8_get_expr: expr LS expr RS
	*/
	expr, err := _parseU8CommonParts(parser)
	if err != nil { return nil, err }

	return asts.NewU8GetAST(expr[0], expr[1]), nil
}
