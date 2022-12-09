package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
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

func a() {

	input := ``
	input = `A Y
B X
C Z`
	input = ""
	list := process(input)
	r := m.Sum(list)
	// r := m.Sum(list)
	fmt.Println(list)
	fmt.Println(r)

}

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func process(s string) []int {
	if s != "" {
		input.SetRaw(s)
	}

	// var things = input.Load()
	var games = input.LoadSliceSliceString("\n", " ")

	list := []int{}
	for _, v := range games {
		score := 0
		switch {
		case v[0] == "A" && v[1] == "X":
			score = 1 + 3
		case v[0] == "A" && v[1] == "Y":
			score = 2 + 6
		case v[0] == "A" && v[1] == "Z":
			score = 3 + 0
		case v[0] == "B" && v[1] == "X":
			score = 1 + 0
		case v[0] == "B" && v[1] == "Y":
			score = 2 + 3
		case v[0] == "B" && v[1] == "Z":
			score = 3 + 6
		case v[0] == "C" && v[1] == "X": // scissors rock
			score = 1 + 6
		case v[0] == "C" && v[1] == "Y": // scissors paper
			score = 2 + 0
		case v[0] == "C" && v[1] == "Z": // scissors scissors
			score = 3 + 3
		}
		list = append(list, score)
	}

	return list
}

var (
	_ = m.Abs
)
