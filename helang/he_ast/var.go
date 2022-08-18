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

func (ast *VarDefAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, err := ast.val.evaluate(env)
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

func (ast *VarAssignAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	_, ok := env[ast.ident]
	if !ok {
		return core.NewU8Empty(), fmt.Errorf("CyberNameException: %s is not defined", ast.ident)
	}
	val, err := ast.val.evaluate(env)
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

func (ast *VarIncrementAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, ok := env[ast.ident]
	if ok {
		val.Increment()
	} else {
		return core.NewU8Empty(), fmt.Errorf("CyberNameException: %s is not defined", ast.ident)
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

func (ast *VarExprAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, ok := env[ast.ident]
	if ok {
		return val, nil
	} else {
		return core.NewU8Empty(), fmt.Errorf("CyberNameException: %s is not defined", ast.ident)
	}
}