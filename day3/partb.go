package main

import (
	"fmt"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

func b() {

	s := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	s = ""

	input := loadInput(s)
	list := processB(input)

	fmt.Println(string(list))
	t := 0
	for _, v := range list {
		t += score(v)
	}
	fmt.Println(t)
}

func processB(things []string) []rune {

	list := []rune{}
	for i := 0; i < len(things); i += 3 {
		a := map[rune]bool{}
		b := map[rune]bool{}
		c := map[rune]bool{}

		for _, v := range things[i] {
			a[v] = true
		}
		for _, v := range things[i+1] {
			b[v] = true
		}
		for _, v := range things[i+2] {
			c[v] = true
		}

		// find the letter that is in both a and b
		for k := range a {
			if b[k] && c[k] {
				list = append(list, k)
				break
			}
		}
	}
	return list
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
