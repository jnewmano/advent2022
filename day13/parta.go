package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

// 4809

// To run, use ./r a f
// change f to t to use example input file
// a to b to run part b
func a() {
	s := ``
	s = ``

	things := loadInput(s)
	list := process(things)

	indexes := []int{}
	for i, v := range list {
		if v {
			indexes = append(indexes, i+1)
		}
	}
	fmt.Println(indexes)
	fmt.Println(m.Sum(indexes))
}

func process(pairs [][]string) []bool {

	list := []bool{}
	for _, v := range pairs {
		var a, b any
		err := json.Unmarshal([]byte(v[0]), &a)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(v[1]), &b)
		if err != nil {
			log.Fatal(err)
		}

		s := compare(a, b)
		list = append(list, s == 1)
	}

	return list
}

func compare(a, b any) int {

	// all inputs are slices
	left := a.([]any)
	right := b.([]any)

	for i := range left {
		if i >= len(right) {
			// right has run out of items
			return -1
		}

		leftInt := isInt(left[i])
		rightInt := isInt(right[i])

		if leftInt && rightInt {
			if m.Int(left[i]) < m.Int(right[i]) {
				return 1
			}
			if m.Int(left[i]) > m.Int(right[i]) {
				return -1
			}
			// if equal, keep going
		} else if !leftInt && !rightInt {
			s := compare(left[i], right[i])
			if s != 0 {
				return s
			}
		} else {
			var s int
			if leftInt {
				s = compare([]any{left[i]}, right[i])
			} else {
				s = compare(left[i], []any{right[i]})
			}
			if s != 0 {
				return s
			}
		}
	}

	// we've processed all of the items in the left list
	if len(left) < len(right) {
		// If the left list runs out of items first, the inputs are in the right order
		return 1
	}

	// need to keep running checks
	return 0
}

func isInt(v interface{}) bool {
	_, okInt := v.(int)
	_, okFloat64 := v.(float64)
	return okInt || okFloat64
}

func loadInput(s string) [][]string {
	if s != "" {
		input.SetRaw(s)
	}
	return input.LoadSliceSliceString("\n\n", "\n")
}

func main() {
	part := strings.ToLower(os.Getenv("PART"))
	switch part {
	case "", "a":
		fmt.Println("Running part a")
		a()
	case "b":
		fmt.Println("Running part b")
		b()
	default:
		log.Fatalf("unknown part [%s] use either a or b\n", part)
	}

}

var (
	_ = m.Abs
)
