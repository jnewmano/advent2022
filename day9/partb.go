package main

import (
	"fmt"
)

func b() {
	s := ``

	things := loadInput(s)
	visited := process(things, 10)

	fmt.Println(len(visited))
}
