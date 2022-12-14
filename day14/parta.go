package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

const (
	rockValue = 1
	sandValue = 2
)

// To run, use ./r a f
// change f to t to use example input file
// a to b to run part b
func a() {
	s := ``
	s = ``

	paths := loadInput(s)
	grid := rockGrid(paths)

	maxY := 0
	for v := range grid {
		if v.Y > maxY {
			maxY = v.Y
		}
	}

	process(grid, maxY, maxY+2)

	fmt.Println("sand", count(grid))
}

func count(grid map[m.Point]*m.Node) int {
	sand := 0
	for _, v := range grid {
		if v.Value == sandValue {
			sand++
		}
	}
	return sand
}

func process(grid map[m.Point]*m.Node, maxY int, floor int) {

	m.DrawMapValues(grid, 0)

	moves := []m.Point{
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}

	for {
		sand := m.Point{X: 500, Y: 0}
	sand:
		for {
			// let's move the sand
			moved := false
			for _, move := range moves {
				dest := sand.Move(move)
				if dest.Y >= floor {
					break sand
				}

				if n := grid[dest]; n == nil || n.Value == 0 {
					moved = true
					sand = dest
					break
				}
			}

			if sand.Y > maxY {
				break // we are done, sand fell off the map
			}

			if sand.Y == 0 || !moved {
				// no more sand can fall
				break
			}

		}

		grid[sand] = &m.Node{Point: sand, Value: sandValue}

		if sand.Y == 0 || sand.Y > maxY {
			break
		}
	}
}

func rockGrid(paths [][]m.Point) map[m.Point]*m.Node {
	all := map[m.Point]*m.Node{}

	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			// draw the lines between points
			start := path[i]
			next := path[i+1]

			min := m.MinCorner(start, next)
			max := m.MaxCorner(start, next)

			for x := min.X; x <= max.X; x++ {
				for y := min.Y; y <= max.Y; y++ {
					p := m.Point{X: x, Y: y}
					all[p] = &m.Node{Point: p, Value: rockValue}
				}
			}
		}
	}
	return all
}

func loadInput(s string) [][]m.Point {
	if s != "" {
		input.SetRaw(s)
	}

	lines := input.LoadSliceString("")

	paths := [][]m.Point{}
	for _, v := range lines {
		points := []m.Point{}
		parts := strings.Split(v, " -> ")
		for _, v := range parts {
			x, y, _ := strings.Cut(v, ",")
			p := m.Point{X: m.Int(x), Y: m.Int(y)}
			points = append(points, p)
		}
		paths = append(paths, points)
	}
	return paths
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
