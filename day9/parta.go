package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

// To run, use ./r a f
// change f to t to use example input file
// a to b to run part b
func a() {
	s := ``

	things := loadInput(s)

	visited := process(things, 2)

	fmt.Println(len(visited))
}

func process(things [][]string, numberKnots int) map[m.Point]int {
	knots := make([]*m.Point, numberKnots)
	for i := range knots {
		knots[i] = &m.Point{}
	}

	var visited = map[m.Point]int{}
	for _, v := range things {
		step := m.Int(v[1])
		for ; step > 0; step-- {
			// move the head
			h := knots[0]
			switch v[0] {
			case "L":
				h.X -= 1
			case "R":
				h.X += 1
			case "U":
				h.Y += 1
			case "D":
				h.Y -= 1
			default:
				log.Fatalf("unknown direction [%s]\n", v[0])
			}

			for kn := 1; kn < numberKnots; kn++ {
				h := knots[kn-1]
				t := knots[kn]

				// move the tail as needed
				dx := h.X - t.X
				dy := h.Y - t.Y
				if m.Abs(dx)+m.Abs(dy) > 2 {
					// handle diagonal step
					if dx > 0 && dy > 0 {
						t.X++
						t.Y++
					} else if dx > 0 && dy < 0 {
						t.X++
						t.Y--
					} else if dx < 0 && dy > 0 {
						t.X--
						t.Y++
					} else if dx < 0 && dy < 0 {
						t.X--
						t.Y--
					}
				} else if m.Abs(dx) > 1 {
					if dx > 0 {
						t.X++
					} else {
						t.X--
					}
				} else if m.Abs(dy) > 1 {
					if dy > 0 {
						t.Y++
					} else {
						t.Y--
					}
				}

				if kn == numberKnots-1 {
					// track where the last tail knot has been
					visited[*t] = visited[*t] + 1
				}
			}
		}
	}
	return visited
}

func loadInput(s string) [][]string {
	if s != "" {
		input.SetRaw(s)
	}

	return input.LoadSliceSliceString("\n", " ")
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
