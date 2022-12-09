package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

// To run, use ./r a f
// change f to t to use example input file
// a to b to run part b
func a() {
	s := ``
	s = ``

	things := loadInput(s)
	list := process(things)

	fmt.Println("sum of list :", m.Sum(list))
	fmt.Println("high in list:", m.High(list))
	fmt.Println("low in list :", m.Low(list))
	fmt.Println("len of list :", len(list))

}

func process(things []string) []int {

	list := []int{}
	for _, v := range things {
		fmt.Println(v)
	}

	return list
}

func loadInput(s string) []string {
	if s != "" {
		input.SetRaw(s)
	}
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	return input.LoadSliceString("")
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
