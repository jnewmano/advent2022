package simplemath

import (
	"sort"
)

// PopHigh returns the highest number in a slice and removes it from the slice
func PopHigh(list []int) (int, []int) {
	h := High(list)

	return h, Remove(list, h)
}

// High returns the highest number in a slice
func High(in ...any) int {
	scores := IntSlice(in)

	n := Clone(scores...)
	sort.Ints(n)
	if len(n) == 0 {
		return 0
	}

	return n[len(n)-1]
}

// PopLow returns the lowest number in a slice and removes it from the slice
func PopLow(list []int) (int, []int) {
	l := Low(list)

	return l, Remove(list, l)
}

// Low returns the lowest number in a slice
func Low(in ...any) int {
	scores := IntSlice(in)

	n := Clone(scores...)
	sort.Ints(n)
	return n[0]
}

// Sum returns the sum of a slice
func Sum(in ...any) int {
	scores := IntSlice(in)

	sum := 0
	for _, v := range scores {
		sum += v
	}
	return sum
}

// Sum returns the product of a slice
func Product(in ...any) int {
	scores := IntSlice(in)

	sum := 1
	for _, v := range scores {
		sum *= v
	}
	return sum
}

// Clone returns a copy of a slice
func Clone[T any](list ...T) []T {
	n := make([]T, len(list))
	copy(n, list)
	return n
}

// Remove removes a value from a slice
// If a number is present multiple times, it will remove the first one
func Remove[T comparable](list []T, num T) []T {
	for i, v := range list {
		if v == num {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
