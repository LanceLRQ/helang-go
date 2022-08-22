package tests

import (
	assertProvider "github.com/stretchr/testify/assert"
	"helang-go/helang/asts"
	"helang-go/helang/core"
	"testing"
)

type ListTestCase struct {
	a *asts.VarExprAST
	b *asts.VarExprAST
	c *asts.VarExprAST
	env map[string]*core.U8
}

func NewListTestCase() ListTestCase {
	return ListTestCase{
		a: asts.NewVarExprAST("a"),
		b: asts.NewVarExprAST("b"),
		c: asts.NewVarExprAST("c"),
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
	ret, err := asts.NewU8GetAST(case1.a, case1.b).Evaluate(case1.env)
	if err != nil {
		t.Fatal(err)
	}
	expectRet := core.NewU8Array([]int { 1, 3 })
	assert.Equal(expectRet, ret, "value should be '1 | 3'")
}

func TestListSet(t *testing.T) {
	assert := assertProvider.New(t)
	case1 := NewListTestCase()
	_, err := asts.NewU8SetAST(case1.a, case1.b, case1.c).Evaluate(case1.env)
	if err != nil {
		t.Fatal(err)
	}
	expectRet := core.NewU8Array([]int { 12, 2, 12, 4 })
	actualRet := case1.env["a"]
	assert.Equal(expectRet, actualRet, "value should be '12 | 2 | 12 | 4'")
}