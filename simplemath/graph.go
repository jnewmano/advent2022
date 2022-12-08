package simplemath

type Point struct {
	X int
	Y int
}

// default type for map, usually wrapped in another type to track additional information about the node
type Node struct {
	Point Point
	Value int
	Meta  map[string]interface{}

	// fields used for graph algorithms
	Counter int
	Touched bool
}

func PointNeighbors(p Point) []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
	}
}

// Neighbors returns the neighbors of a point
func NodeNeighbors[T any](nodes map[Point]T, p Point) []T {

	neighborPoints := PointNeighbors(p)

	neighbors := []T{}
	for _, v := range neighborPoints {
		n, ok := nodes[v]
		if !ok { // node doesn't exist
			continue
		}

		neighbors = append(neighbors, n)
	}

	return neighbors
}
