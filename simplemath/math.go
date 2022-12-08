package simplemath

import (
	"fmt"
	"strconv"
)

// Int converts a things into an int
func Int(in any) int {
	switch v := in.(type) {
	case uint8: // assume an ascii character
		if v < '0' || v > '9' {
			panic(fmt.Errorf("invalid character %q, unable to convert to int", v))
		}
		return int(v) - '0'
	case int:
		return v
	case int64:
		return int(v)
	case string:
		r, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return r
	case nil:
	default:
		panic(fmt.Errorf("unhandled type %T", in))
	}
	return 0
}

// Abs returns the absolute value of a thing
func Abs(in any) int {
	v := Int(in)

	if v < 0 {
		return -v
	}
	return v
}

// LCM returns the least common multiple
func LCM(a ...int) int {
	l := a[0]
	for i := 1; i < len(a); i++ {
		l = lcm(l, a[i])
	}
	return l
}

func lcm(a, b int) int {
	return a * b / GCD(a, b)
}

// GCD returns the greatest common denominator of
// a list of numbers
func GCD(a ...int) int {
	g := a[0]
	for i := 1; i < len(a); i++ {
		g = gcd(g, a[i])
	}
	return g
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}
