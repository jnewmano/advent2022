package main

import (
	"fmt"
)

func b() {
	s := ``
	things := loadInput(s)
	history := process(things)

	for i, v := range history[:40*6] {
		char := " "
		pos := (i % 40)
		if v == pos || v == pos+1 || v == pos-1 {
			char = "#"
		}
		fmt.Print(char)
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}
