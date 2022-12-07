package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
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
		log.Fatalln("unknown part", part)
	}
}

func a() {
	list := process()
	fmt.Println(m.High(list))
}

func process() []int {
	var things = input.LoadSliceSliceString("\n\n", "\n")

	var list []int
	for _, v := range things {
		sum := m.Sum(v)
		list = append(list, sum)
	}
	return list
}
