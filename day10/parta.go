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
	things := loadInput(s)
	history := process(things)

	e := []int{
		20, 60, 100, 140, 180, 220,
	}
	score := 0
	for _, v := range e {
		s := history[v-1] * v
		score += s
	}
	fmt.Println("score", score)
}

func process(lines []string) []int {

	register := 1

	pc := 0

	history := make([]int, len(lines)*2)
	for _, v := range lines {
		parts := strings.Split(v, " ")
		instruction := parts[0]

		switch instruction {
		case "addx":
			history[pc] = register
			history[pc+1] = register
			val := m.Int(parts[1])
			register += val
			pc += 2
		case "noop":
			history[pc] = register
			pc++
		default:
			log.Fatalln("unknown instruction", instruction)
		}
	}
	history[pc] = register

	return history
}

func loadInput(s string) []string {
	if s != "" {
		input.SetRaw(s)
	}
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	return input.LoadSliceString("\n")
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
