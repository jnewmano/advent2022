package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func a() {
	s := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	s = ``

	things := loadInput(s)
	list := process(things)

	// fmt.Println("sum of list :", m.Sum(list))
	// fmt.Println("high in list:", m.High(list))
	// fmt.Println("low in list :", m.Low(list))
	fmt.Println("len of list :", len(list))
}

func loadInput(s string) [][]string {
	if s != "" {
		input.SetRaw(s)
	}
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	return input.LoadSliceSliceString("", ",")
}

func process(things [][]string) []int {

	list := []int{}
	for i, v := range things {
		a := strings.Split(v[0], "-")
		b := strings.Split(v[1], "-")
		a1 := m.Int(a[0])
		a2 := m.Int(a[1])

		b1 := m.Int(b[0])
		b2 := m.Int(b[1])

		if a1 >= b1 && a2 <= b2 {
			fmt.Println("a1 >= b1 && a2 <= b2", i, v, a1, b1, a1 >= b1)
			list = append(list, i)
			continue
		}
		if b1 >= a1 && b2 <= a2 {
			fmt.Println("other", i, v)
			list = append(list, i)
			continue
		}
	}

	return list
}

var (
	_ = m.Abs
)

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
