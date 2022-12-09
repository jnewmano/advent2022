package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

// 3153

func a() {

	s := ``

	things := loadInput(s)
	list := process(things)

	count := 0
	for _, v := range list {

		for i := 0; i < len(list); i++ {
			vv := v[i]
			if vv {
				count++
			}
		}
	}

	fmt.Println(count)
}

func loadInput(s string) []string {
	if s != "" {
		input.SetRaw(s)
	}
	return input.LoadSliceString("")
}

func process(things []string) [][]bool {

	see := m.New2DSlice[bool](len(things), len(things[0]))

	for i := 0; i < len(things); i++ {
		maxH := -1
		for j := 0; j < len(things); j++ {
			// check visible from the left
			if h(things, i, j) > maxH {

				see[i][j] = true
				maxH = h(things, i, j)
			}
		}
	}

	for i := 0; i < len(things); i++ {
		maxH := -1
		for j := len(things) - 1; j >= 0; j-- {
			// check visible from the right
			if h(things, i, j) > maxH {
				see[i][j] = true
				maxH = h(things, i, j)
			}
		}
	}

	for j := 0; j < len(things); j++ {
		maxH := -1
		for i := 0; i < len(things); i++ {
			// check visible from the left
			if h(things, i, j) > maxH {
				see[i][j] = true
				maxH = h(things, i, j)
			}
		}
	}

	for j := 0; j < len(things); j++ {
		maxH := -1
		for i := len(things) - 1; i >= 0; i-- {
			// check visible from the right
			if h(things, i, j) > maxH {
				see[i][j] = true
				maxH = h(things, i, j)
			}
		}
	}

	return see
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

func h(things []string, x int, y int) int {
	return m.Int(things[x][y])
}

var (
	_ = m.Abs
)
