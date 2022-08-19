package he_ast

import (
	"fmt"
	"helang-go/helang/core"
)

// ===  VarDefAST ===

type VarDefAST struct {
	BasicAST
	ident string
	val   BasicAST
}

func NewVarDefAST (ident string, val BasicAST) *VarDefAST {
	return &VarDefAST{ ident: ident, val: val }
}

func (ast *VarDefAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, err := ast.val.Evaluate(env)
	if err != nil {
		return val, err
	}
	env[ast.ident] = val
	return core.NewU8Empty(), nil
}

// ===  VarAssignAST ===

type VarAssignAST struct {
	BasicAST
	ident string
	val   BasicAST
}

func NewVarAssignAST (ident string, val BasicAST) *VarAssignAST {
	return &VarAssignAST{ ident: ident, val: val }
}

func (ast *VarAssignAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	_, ok := env[ast.ident]
	if !ok {
		return core.NewU8Empty(), fmt.Errorf("%w: %s is not defined", core.CyberNameException, ast.ident)
	}
	val, err := ast.val.Evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	env[ast.ident] = val
	return val, nil
}

// === VarIncrementAST ===

type VarIncrementAST struct {
	BasicAST
	ident string
}

func NewVarIncrementAST (ident string) *VarIncrementAST {
	return &VarIncrementAST{ ident: ident }
}

func (ast *VarIncrementAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, ok := env[ast.ident]
	if ok {
		val.Increment()
	} else {
		return core.NewU8Empty(), fmt.Errorf("%w: %s is not defined", core.CyberNameException, ast.ident)
	}
	return val, nil
}

// === VarExprAST ===

type VarExprAST struct {
	BasicAST
	ident string
}

func NewVarExprAST (ident string) *VarExprAST {
	return &VarExprAST{ ident: ident }
}

func (ast *VarExprAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, ok := env[ast.ident]
	if ok {
		return val, nil
	} else {
		return core.NewU8Empty(), fmt.Errorf("%w: %s is not defined", core.CyberNameException, ast.ident)
	}
}

// === SubtractionAST ===

type SubtractionAST struct {
	BasicAST
	first BasicAST
	second BasicAST
}

func NewSubtractionAST (first, second BasicAST) *SubtractionAST {
	return &SubtractionAST{ first: first, second: second }
}

func (ast *SubtractionAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	first, err := ast.first.Evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	second, err := ast.second.Evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	return first.Sub(second)
}

// === ListAST ===

type ListAST struct {
	BasicAST
	asts []BasicAST
}

func NewListAST (asts []BasicAST) *ListAST {
	return &ListAST{ asts: asts }
}

func (ast *ListAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	for _, v := range ast.asts {
		_, err := v.Evaluate(env)
		if err != nil {
			return core.NewU8Empty(), err
		}
	}
	return core.NewU8Empty(), nil
}

// === VoidAST ===

type VoidAST struct {
	BasicAST
}

func NewVoidAST () *VoidAST {
	return &VoidAST{}
}

func (ast *VoidAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	return core.NewU8Empty(), nil
}

