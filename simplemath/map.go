package simplemath

import (
	"fmt"
)

func DrawMapValues(nodes map[Point]*Node, width int) {

	if width == 0 {
		max := 0
		for _, v := range nodes {
			if v.Value > max {
				max = v.Value
			}
		}
		switch {
		case max < 10:
			width = 2
		case max < 100:
			width = 3
		default:
			width = 4
		}
	}

	fmtString := fmt.Sprintf("%%%dd", width)
	f := func(n *Node) string {
		v := 0
		if n != nil {
			v = n.Value
		}
		return fmt.Sprintf(fmtString, v)
	}
	DrawMap(nodes, f)
}

func DrawMap[T any](nodes map[Point]T, s func(T) string) {
	// draw the map and the points we touched

	minY := -1
	minX := -1
	maxY := 0
	maxX := 0
	for v := range nodes {
		if v.Y < minY || minY == -1 {
			minY = v.Y
		}
		if v.X < minX || minX == -1 {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Point{X: x, Y: y}
			n := nodes[p]
			fmt.Print(s(n))
		}
		fmt.Print("\n")
	}
}
