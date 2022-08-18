package he_ast

import (
	"helang-go/helang/core"
)

// === EmptyU8InitAST ===

type EmptyU8InitAST struct {
	BasicAST
	length int
}

func NewEmptyU8InitAST (length int) *EmptyU8InitAST {
	return &EmptyU8InitAST{ length: length }
}

func (ast *EmptyU8InitAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	return core.NewU8Array(make([]int, ast.length, 0)), nil
}

// === OrU8InitAST ===

type OrU8InitAST struct {
	// How the King He defines uint8 list: by | operator.
	BasicAST
	first int
	second *OrU8InitAST
}

func NewOrU8InitAST(first int, second *OrU8InitAST) *OrU8InitAST {
	return &OrU8InitAST{ first: first, second: second }
}

func (ast *OrU8InitAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	if ast.second == nil {
		return core.NewU8Array([]int{ ast.first }), nil
	}

	second, err := ast.second.evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	elements := []int{ ast.first }
	elements = append(elements, second.Value...)
	return core.NewU8Array(elements), nil
}

// === EmptyU8InitAST ===

type U8SetAST struct {
	BasicAST
	list BasicAST
	subscript BasicAST
	value BasicAST
}

func NewU8SetAST (list, subscript, value BasicAST) *U8SetAST {
	return &U8SetAST{ list: list, subscript: subscript, value: value}
}

func (ast *U8SetAST) evaluate(env map[string]*core.U8) (*core.U8, error) {
	lst, err := ast.list.evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	subscripts, err := ast.subscript.evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}
	val, err := ast.value.evaluate(env)
	if err != nil {
		return core.NewU8Empty(), err
	}

	// TODO SETITEM

	return core.NewU8Empty(), nil
}
