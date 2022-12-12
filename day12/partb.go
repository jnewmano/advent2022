package main

import (
	"fmt"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

func b() {
	s := ""

	allNodes, _, _ := loadInput(s)

	// dynamic start point, find the shortest path to the end point
	ls := []int{}
	for p, n := range allNodes {
		if n.Value != 'a' {
			// only test valid start points
			continue
		}
		nodes, _, endPoint := loadInput(s)
		start := nodes[p]
		end := nodes[endPoint]

		process(nodes, start, end)
		if start.Counter == 0 {
			// no path found, skip it
			continue
		}
		ls = append(ls, start.Counter)
	}

	// the cost counter is the same as the number of steps
	// return the lowest step count found
	fmt.Println(m.Low(ls))
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
