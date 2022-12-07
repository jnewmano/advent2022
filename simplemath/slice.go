package simplemath

import (
	"fmt"
)

// IntSlice converts of slice of anything into a slice of ints
func IntSlice(in any) []int {

	switch v := in.(type) {
	case int:
		return []int{v}
	case string:
		return []int{Int(v)}
	case []int:
		return v
	case []string:
		b := make([]int, len(v))
		for i, x := range v {
			b[i] = Int(x)
		}
		return b
	case []interface{}:
		b := []int{}
		for _, x := range v {
			b = append(b, IntSlice(x)...)
		}
		return b
	case nil:
		return []int{}
	default:
		panic(fmt.Errorf("unhandled type %T", in))
	}
}
