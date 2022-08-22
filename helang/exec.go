package helang

import "helang-go/helang/core"

func RunCode(code string, env map[string]*core.U8) (*core.U8, error) {
	lex := NewLex([]byte(code))
	tokens, err := lex.Lex()
	if err != nil {
		return nil, err
	}
	parser := NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	ret, err := ast.Evaluate(env)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func RunCodeWithoutEnv(code string) (map[string]*core.U8, *core.U8, error) {
	env := map[string]*core.U8{}
	ast, err := RunCode(code, env)
	return env, ast, err
}
