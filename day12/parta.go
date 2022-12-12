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

	s := ""

	nodes, startPoint, endPoint := loadInput(s)
	start := nodes[startPoint]
	end := nodes[endPoint]

	process(nodes, start, end)

	// the cost counter is the same as the number of steps
	fmt.Println(start.Counter)
}

type MapType map[m.Point]*m.Node

func process(nodes MapType, start *m.Node, end *m.Node) {

	// initialize the queue with the end node
	queue := []*m.Node{
		end,
	}

	// go through every element in the queue, including new elements added
	// to the end over the course of the algorithm, and for each element,
	// do the following
	pos := 0
	for {

		if pos >= len(queue) {
			// break when we've reached the end of the queue
			// we can't end when we reach the starting point because we're
			// not using a priority queue, so we have to search paths
			// until we're no longer adding paths
			break
		}

		// grab the next node off of the queue
		c := queue[pos]
		pos++

		// Get the neighbors and determine the cost to get to the next neighbor
		neighbors := m.NodeNeighbors(nodes, c.Point)
		for _, n := range neighbors {
			if n.Point == end.Point {
				// end point, don't update
				continue
			}

			// condition for allowing the transition
			diff := c.Value - n.Value
			if diff > 1 {
				// we can't move up more than one step, we can move down any number of steps
				continue
			}

			// cost calculation of the transition
			cost := 1 // cost is always 1, if the transition is allowed
			newCounter := c.Counter + cost

			if n.Counter != 0 && newCounter >= n.Counter {
				// we've already found a shorter path to this node, skip it
				continue
			}

			// update the counter with the accumulated cost and append the item to the queue
			n.Counter = newCounter
			queue = append(queue, n)
		}
	}
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

type Node m.Node

func loadInput(s string) (MapType, m.Point, m.Point) {
	if s != "" {
		input.SetRaw(s)
	}

	nodes := make(MapType)
	lines := input.LoadSliceString("")

	var start m.Point
	var end m.Point

	for y, row := range lines {
		for x, v := range []byte(row) {
			p := m.Point{X: x, Y: y}
			if v == 'S' {
				start = p
				v = 'a'
			}
			if v == 'E' {
				end = p
				v = 'z'
			}

			n := m.Node{
				Point: p,
				Value: int(v),
			}

			nodes[p] = &n
		}
	}

	return nodes, start, end
}
