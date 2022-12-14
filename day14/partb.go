package main

import (
	"fmt"
)

func b() {
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

	process(grid, maxY+5, maxY+2)

	fmt.Println("sand", count(grid))
}
