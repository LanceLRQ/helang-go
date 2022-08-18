package core

import (
	"strconv"
	"strings"
)

type U8 struct {
	Value []int
}

func NewU8Empty () *U8 {
	return &U8{
		Value: []int{},
	}
}
func NewU8Int (val int) *U8 {
	return &U8{
		Value: []int{ val },
	}
}

func NewU8Array (val []int) *U8 {
	return &U8{
		Value: val,
	}
}

func (u8 *U8) ToString () string {
	outputStr := make([]string, 0, 10)
	for _, v := range u8.Value {
		outputStr = append(outputStr, strconv.Itoa(v))
	}
	return strings.Join(outputStr, " | ")
}

func (u8 *U8) Equal (other *U8) bool {
	return other.ToString() == u8.ToString()
}

func (u8 *U8) EqualIntArray (other []int) bool {
	if len(other) != len(u8.Value) {
		return false
	}
	for i, o := range other {
		if u8.Value[i] != o {
			return false
		}
	}
	return true
}

func (u8 *U8) Increment () {
	for i, _ := range u8.Value {
		u8.Value[i]++
	}
}