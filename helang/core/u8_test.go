package core

import (
	assertProvider "github.com/stretchr/testify/assert"
	"testing"
)

func TestU8Compare(t *testing.T) {
	assert := assertProvider.New(t)
	a := NewU8Array([]int{1, 2 ,3})
	b := NewU8Array([]int{1 ,2, 4})
	assert.Equal(a, a)
	assert.NotEqual(a, b)
}

func TestU8Subtraction(t *testing.T) {
	assert := assertProvider.New(t)
	a := NewU8Array([]int{6, 5 ,4})
	b := NewU8Array([]int{1 ,2, 4})
	c := NewU8Int(2)

	aSb, err := a.Sub(b)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(NewU8Array([]int{5, 3, 0}), aSb)

	bSa, err := b.Sub(a)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(NewU8Array([]int{-5, -3, 0}), bSa)

	aSc, err := a.Sub(c)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(NewU8Array([]int{4, 3, 2}), aSc)

	bSc, err := b.Sub(c)
	if err != nil { t.Fatal(err) }
	assert.EqualValues(NewU8Array([]int{-1, 0, 2}), bSc)

	_, err = c.Sub(a)
	assert.ErrorIs(err, CyberArithmeticException)
}

func TestU8SetAll(t *testing.T) {
	assert := assertProvider.New(t)
	a := NewU8Array([]int{1, 2, 3})

	err := a.SetItem(NewU8Int(0), NewU8Int(10))
	if err != nil { t.Fatal(err) }

	assert.EqualValues(NewU8Array([]int{10, 10, 10}), a)
}