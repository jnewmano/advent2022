package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

func a() {
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
	list := process(stacks, things)

	/*	fmt.Println("sum of list :", m.Sum(list))
		fmt.Println("high in list:", m.High(list))
		fmt.Println("low in list :", m.Low(list))
		fmt.Println("len of list :", len(list))
	*/

	for i, v := range list {
		if i == 0 {
			continue
		}
		fmt.Print(v[0])
	}
	fmt.Println()
}

func loadInput(s string) ([10][]byte, []string) {
	if s != "" {
		input.SetRaw(s)
	}

	stacks := [10][]byte{}
	stacks[1] = []byte("FLMW")
	stacks[2] = []byte("FMVZB")
	stacks[3] = []byte("QLSRVH")
	stacks[4] = []byte("JTMPQVSF")
	stacks[5] = []byte("WSL")
	stacks[6] = []byte("WJRMPVF")
	stacks[7] = []byte("FRNPCQJ")
	stacks[8] = []byte("BRWZSPHV")
	stacks[9] = []byte("WZHGCJMB")
	/*
	               [J]             [B] [W]
	               [T]     [W] [F] [R] [Z]
	           [Q] [M]     [J] [R] [W] [H]
	       [F] [L] [P]     [R] [N] [Z] [G]
	   [F] [M] [S] [Q]     [M] [P] [S] [C]
	   [L] [V] [R] [V] [W] [P] [C] [P] [J]
	   [M] [Z] [V] [S] [S] [V] [Q] [H] [M]
	   [W] [B] [H] [F] [L] [F] [J] [V] [B]
	    1   2   3   4   5   6   7   8   9
	*/
	// var things = input.Load()
	var things = input.LoadSliceString("\n\n")
	instructions := strings.Split(things[1], "\n")
	return stacks, instructions
}

func process(stacks [10][]byte, instructions []string) [10][]byte {

	for _, v := range instructions {
		parts := strings.Split(v, " ")
		num := m.Int(parts[1])
		from := m.Int(parts[3])
		to := m.Int(parts[5])

		for i := 0; i < num; i++ {
			stacks[to] = append([]byte{stacks[from][0]}, stacks[to]...)
			stacks[from] = stacks[from][1:]
		}
		/*
			take := stacks[from][len(stacks[from])-num:]
			fmt.Println("moving", num, take, " to ", to)
			stacks[from] = stacks[from][0 : len(stacks[from])-num]
			stacks[to] = append(stacks[to], take...)
		*/
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
