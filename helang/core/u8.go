package core

import (
	"fmt"
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


func (u8 *U8) Sub(other *U8) (*U8, error) {
	if len(other.Value) == 1 || len(other.Value) == len(u8.Value) {
		// Normal subtraction && Vector subtraction.
		arr := make([]int, 0, 1)
		for _, ov := range other.Value {
			for _, v := range u8.Value {
				arr = append(arr, v - ov)
			}
		}
		return NewU8Array(arr), nil
	}
	return NewU8Empty(), fmt.Errorf("%w: illegal subtraction: %s - %s", CyberSubtractionException, u8.ToString(), other.ToString())
}


func (u8 *U8) SetItem(subscripts *U8, value *U8) error {
	if len(value.Value) > 1 {
		return fmt.Errorf("%w: no high dimension u8", CyberNotSupportedException)
	}
	if len(value.Value) == 0 {
		return fmt.Errorf("%w: you must set u8 with single value", CyberNotSupportedException)
	}
	val := value.Value[0]

	// Set all elements if subscript is single 0.
	if subscripts.EqualIntArray([]int{0}) {
		for i, _ := range u8.Value {
			u8.Value[i] = val
		}
		return nil
	}

	for _, subscript := range subscripts.Value {
		if subscript == 0 {
			return fmt.Errorf("%w: subscript 0 is designed for setting all elements you should write like array[0] = 10", CyberNotSupportedException)
		}
		u8.Value[subscript - 1] = val
	}
	return nil
}

func utilValueInArray(array []int, value int) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}

func (u8 *U8) GetItem(subscripts *U8) *U8 {
	// Like the operation of sublist.
	// And Saint He likes arrays whose subscript start from 1.

	arr := make([]int, 0, 1)

	for i := 1; i < len(u8.Value) + 1; i++ {
		if utilValueInArray(subscripts.Value, i) {
			arr = append(arr, u8.Value[i - 1])
		}
	}
	return NewU8Array(arr)
}