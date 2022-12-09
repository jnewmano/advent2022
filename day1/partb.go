package main

import (
	"fmt"

	m "github.com/jnewmano/advent2022/simplemath"
)

func b() {
	list := process()
	h1, list := m.PopHigh(list)
	h2, list := m.PopHigh(list)
	h3, list := m.PopHigh(list)
	fmt.Println(h1 + h2 + h3)
}
