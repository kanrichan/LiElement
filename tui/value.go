package tui

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NewValue("123").Get() == NewValue("123").Get())
}

// BaseValue BaseValue
type BaseValue[T any] struct {
	x T
}

// Set Set
func (v *BaseValue[T]) Set(x T) {
	v.x = x
}

// Get Get
func (v *BaseValue[T]) Get() T {
	return v.x
}

// NewValue NewValue
func NewValue[T any](x T) *BaseValue[T] {
	return &BaseValue[T]{x}
}

type length struct {
	t bool
	l uint64
	f uint64
	p uint64
}

// IsPercentLength IsPercentLength
func IsPercentLength(v *BaseValue[string]) bool {
	if v != nil && strings.Contains(v.Get(), "%") {
		_, err := strconv.ParseUint(
			strings.TrimSuffix(v.Get(), "%"), 10, 64)
		return err == nil
	}
	return false
}

// MeasureLength MeasureLength
func MeasureLength(v *BaseValue[string], base uint64) uint64 {
	if v == nil {
		return 0
	}
	switch {
	case IsPercentLength(v):
		n, _ := strconv.ParseUint(
			strings.TrimSuffix(v.Get(), "%"), 10, 64)
		return (n / 100) * base
	default:
		n, err := strconv.ParseUint(v.Get(), 10, 64)
		if err != nil {
			return 0
		}
		return n
	}
}

func uint64add(x, y uint64) uint64 {
	if math.MaxUint64-x > y {
		return x + y
	}
	return math.MaxUint64
}

func uint64sub(x, y uint64) uint64 {
	if x > y {
		return x - y
	}
	return 0
}

func uint64div(x, y uint64) uint64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func uint64mul(x, y uint64) uint64 {
	if x == 0 || y == 0 {
		return 0
	}
	if math.MaxUint64/x > y {
		return x * y
	}
	return math.MaxUint64
}
