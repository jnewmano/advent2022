package simplemath

type Point struct {
	X int
	Y int
}

// default type for map, usually wrapped in another type to track additional information about the node
type Node struct {
	Point Point
	Value int

	// internal fields
	Counter int
	Touched bool
}

func PointNeighbors(p Point, includeDiagonals ...bool) []Point {
	points := []Point{
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
	}

	if len(includeDiagonals) > 0 && includeDiagonals[0] {
		points = append(points, []Point{
			{X: p.X - 1, Y: p.Y - 1},
			{X: p.X - 1, Y: p.Y + 1},
			{X: p.X + 1, Y: p.Y - 1},
			{X: p.X + 1, Y: p.Y + 1},
		}...)
	}
	return points
}

// NodeNeighbors returns the neighbors of a point
func NodeNeighbors[T any](nodes map[Point]T, p Point, includeDiagonals ...bool) []T {

	neighborPoints := PointNeighbors(p, includeDiagonals...)

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

// GridNeighbors returns the neighbors of a point
func GridNeighbors[T any](grid [][]T, p Point, includeDiagonals ...bool) []Point {

	points := PointNeighbors(p, includeDiagonals...)

	list := []Point{}
	for _, p := range points {
		// filter out the neighors, point neighbors doesn't constrain the neighbors to a grid
		if p.X < 0 || p.X >= len(grid[0]) || p.Y < 0 || p.Y >= len(grid) {
			continue
		}
		list = append(list, p)
	}

	return list
}

func MinPoint(points []*Point) Point {
	x := Min(points, func(v *Point) int { return v.X })
	y := Min(points, func(v *Point) int { return v.Y })
	return Point{X: x, Y: y}
}

func MaxPoint(points []*Point) Point {
	x := Max(points, func(v *Point) int { return v.X })
	y := Max(points, func(v *Point) int { return v.Y })
	return Point{X: x, Y: y}
}
