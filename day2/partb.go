package main

import (
	"fmt"
	"log"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

func b() {
	input := `A Y
B X
C Z`
	input = ""
	list := processB(input)

	r := m.Sum(list)
	// r := m.Sum(list)

	fmt.Println(r)

}

func processB(s string) []int {
	if s != "" {
		input.SetRaw(s)
	}

	// var things = input.Load()
	var games = input.LoadSliceSliceString("\n", " ")

	for i, v := range games {
		// remap the inputs
		// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
		switch {
		case v[0] == "A" && v[1] == "X":
			games[i] = []string{"A", "Z"}
		case v[0] == "A" && v[1] == "Y":
			games[i] = []string{"A", "X"}
		case v[0] == "A" && v[1] == "Z":
			games[i] = []string{"A", "Y"}
		case v[0] == "B" && v[1] == "X":
			games[i] = []string{"B", "X"} // lose paper covers rock
		case v[0] == "B" && v[1] == "Y":
			games[i] = []string{"B", "Y"} // draw paper covers rock
		case v[0] == "B" && v[1] == "Z":
			games[i] = []string{"B", "Z"} // win paper scissors
		case v[0] == "C" && v[1] == "X": // scissors rock
			games[i] = []string{"C", "Y"} // lose scissors, paper
		case v[0] == "C" && v[1] == "Y": // scissors paper
			games[i] = []string{"C", "Z"} // draw scissors, scissors
		case v[0] == "C" && v[1] == "Z": // scissors scissors
			games[i] = []string{"C", "X"} // draw scissors, scissors
		default:
			log.Fatalln("not handled", i)
		}
	}

	list := []int{}
	for i, v := range games {
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
		default:
			log.Fatalln("not handled", i)
		}
		list = append(list, score)
	}

	return list
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
