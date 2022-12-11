package main

import (
	"fmt"
)

func b() {
	s := ``

	monkeys := loadInput(s)
	for i := 0; i < 10000; i++ {
		processRound(monkeys, 1)
	}

	fmt.Println(score(monkeys))
}
