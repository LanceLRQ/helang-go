package asts

import (
	assertProvider "github.com/stretchr/testify/assert"
	"helang-go/helang/core"
	"testing"
)

type ListTestCase struct {
	a *VarExprAST
	b *VarExprAST
	c *VarExprAST
	env map[string]*core.U8
}

func NewListTestCase() ListTestCase {
	return ListTestCase{
		a: NewVarExprAST("a"),
		b: NewVarExprAST("b"),
		c: NewVarExprAST("c"),
		env: map[string]*core.U8{
			"a": core.NewU8Array([]int{ 1, 2, 3, 4 }),
			"b": core.NewU8Array([]int{ 1, 3 }),
			"c": core.NewU8Int(12),
		},
	}
}

func TestListGet(t *testing.T) {
	assert := assertProvider.New(t)
	case1 := NewListTestCase()
	ret, err := NewU8GetAST(case1.a, case1.b).Evaluate(case1.env)
	if err != nil {
		t.Fatal(err)
	}
	expectRet := core.NewU8Array([]int { 1, 3 })
	assert.Equal(expectRet, ret, "value should be '1 | 3'")
}

func TestListSet(t *testing.T) {
	assert := assertProvider.New(t)
	case1 := NewListTestCase()
	_, err := NewU8SetAST(case1.a, case1.b, case1.c).Evaluate(case1.env)
	if err != nil {
		t.Fatal(err)
	}
	expectRet := core.NewU8Array([]int { 12, 2, 12, 4 })
	actualRet := case1.env["a"]
	assert.Equal(expectRet, actualRet, "value should be '12 | 2 | 12 | 4'")
}