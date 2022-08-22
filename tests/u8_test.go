package tests

import (
	assertProvider "github.com/stretchr/testify/assert"
	"helang-go/helang/core"
	"testing"
)

func TestU8Compare(t *testing.T) {
	assert := assertProvider.New(t)
	a := core.NewU8Array([]int{1, 2 ,3})
	b := core.NewU8Array([]int{1 ,2, 4})
	assert.Equal(a, a)
	assert.NotEqual(a, b)
}

func TestU8Subtraction(t *testing.T) {
	assert := assertProvider.New(t)
	a := core.NewU8Array([]int{6, 5 ,4})
	b := core.NewU8Array([]int{1 ,2, 4})
	c := core.NewU8Int(2)

	aSb, err := a.Sub(b)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(core.NewU8Array([]int{ 5, 3, 0}), aSb)

	bSa, err := b.Sub(a)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(core.NewU8Array([]int{ -5, -3, 0}), bSa)

	aSc, err := a.Sub(c)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(core.NewU8Array([]int{ 4, 3, 2}), aSc)

	bSc, err := b.Sub(c)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(core.NewU8Array([]int{ -1, 0, 2}), bSc)

	_, err = c.Sub(a)
	assert.ErrorIs(err, core.CyberArithmeticException)
}

func TestU8SetAll(t *testing.T) {
	assert := assertProvider.New(t)
	a := core.NewU8Array([]int{1, 2, 3})

	err := a.SetItem(core.NewU8Int(0), core.NewU8Int(10))
	if err != nil { t.Fatal(err) }

	assert.EqualValues(core.NewU8Array([]int{ 10, 10, 10}), a)
}