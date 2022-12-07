package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func a() {
	s := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	s = ""

	input := loadInput(s)
	list := process(input)
	fmt.Println(string(list))
	//r := m.High(list)
	// r := m.Sum(list)

	//	fmt.Println(r)
	t := 0
	for _, v := range list {
		t += score(v)
	}
	fmt.Println(t)
}

func score(a rune) int {
	if a >= 'a' && a <= 'z' {
		return int(a - 'a' + 1)
	}

	return int(a - 'A' + 27)

}

func loadInput(s string) []string {
	if s != "" {
		input.SetRaw(s)
	}
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	return input.LoadSliceString("")
}

func process(things []string) []rune {

	list := []rune{}
	for _, line := range things {
		a := map[rune]bool{}
		b := map[rune]bool{}
		for i, v := range line {
			if i < len(line)/2 {
				a[v] = true
			} else {
				b[v] = true
			}
		}
		// find the letter that is in both a and b
		for k := range a {
			if b[k] {
				list = append(list, k)
				break
			}
		}
	}
	return list
}

var (
	_ = m.Abs
)

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
