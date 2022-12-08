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

// BinStringToInt converts a binary string to an integer
// example: "101" -> 5
func BinStringToInt(s []byte) int {
	n := 0
	for _, v := range s {
		n = n<<1 + Int(v)
	}
	return n
}
