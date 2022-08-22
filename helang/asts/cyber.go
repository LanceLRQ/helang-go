package asts

import (
	"fmt"
	"helang-go/helang/core"
	"helang-go/helang/plugins"
	"math/rand"
)

// === PrintAST ===

type PrintAST struct {
	BasicAST
	expr BasicAST
}

func NewPrintAST (expr BasicAST) *PrintAST {
	return &PrintAST{ expr: expr }
}

func (ast *PrintAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	val, err := ast.expr.Evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	fmt.Println(val.ToString())
	return val, nil
}

// === SprintAST ===

type SprintAST struct {
	BasicAST
	expr BasicAST
}

func NewSprintAST (expr BasicAST) *SprintAST {
	return &SprintAST{ expr: expr }
}

func (ast *SprintAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	chars, err := ast.expr.Evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	val := ""
	for _, char := range chars.Value {
		val += string(rune(char))
	}
	fmt.Println(val)
	return chars, nil
}

// === Test5GAST ===

type Test5GAST struct {
	BasicAST
	specialValue []int
}

func NewTest5GAST () *Test5GAST {
	// To avoid coincidence.
	return &Test5GAST{
		specialValue: []int{ rand.Intn(100) + 1, rand.Intn(100) + 1 },
	}
}

func (ast *Test5GAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	plugins.RunSpeedTest()
	return core.NewU8Array(ast.specialValue), nil
}


// === CyberspacesAST ===

type CyberspacesAST struct {
	BasicAST
}

func NewCyberspacesAST () *CyberspacesAST {
	return &CyberspacesAST{}
}

func (ast *CyberspacesAST) Evaluate(env map[string]*core.U8) (*core.U8, error) {
	err := plugins.CheckCyberSpaces()
	if err != nil {
		return core.NewU8Empty(), err
	}
	return core.NewU8Empty(), nil
}


