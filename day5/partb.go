package main

import (
	"fmt"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func b() {
	s := `   [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	s = ``

	stacks, things := loadInput(s)
	/*
		stacks[1] = []string{"N", "Z"}
		stacks[2] = []string{"D", "C", "M"}
		stacks[3] = []string{"P"}
	*/
	list := processB(stacks, things)

	for i, v := range list {
		if i == 0 {
			continue
		}
		fmt.Print(v[0])
	}
	fmt.Println()
}

func processB(stacks [10][]byte, instructions []string) [10][]byte {

	for _, v := range instructions {
		parts := strings.Split(v, " ")
		num := m.Int(parts[1])
		from := m.Int(parts[3])
		to := m.Int(parts[5])

		take := stacks[from][0:num]
		stacks[from] = m.Clone(stacks[from][num:]...)
		stacks[to] = append(take, stacks[to]...)

		for i, v := range stacks {
			if i >= 4 {
				continue
			}
			fmt.Printf("%2d: %v\n", i, v)
		}
		fmt.Println()
	}

	return stacks
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
