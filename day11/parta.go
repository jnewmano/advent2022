package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

// To run, use ./r a f
// change f to t to use example input file
// a to b to run part b
func a() {
	s := ``

	monkeys := loadInput(s)
	for i := 0; i < 20; i++ {
		processRound(monkeys, 3)
	}

	fmt.Println(score(monkeys))
}

func score(monkeys []*Monkey) int {
	counts := []int{}
	for _, v := range monkeys {
		counts = append(counts, v.Count)
	}

	a, counts := m.PopHigh(counts)
	b, _ := m.PopHigh(counts)
	return a * b
}

func processRound(monkeys []*Monkey, scale int) {
	gcf := 1
	for _, monkey := range monkeys {
		gcf *= monkey.Test
	}

	for _, monkey := range monkeys {
		// go through the monkey's items
		for _, item := range monkey.Items {
			monkey.Count++

			// calculate the new worry level
			switch {
			case monkey.Operation[1] == "+":
				item += m.Int(monkey.Operation[2])
			case monkey.Operation[1] == "*" && monkey.Operation[2] == "old":
				item *= item
			case monkey.Operation[1] == "*":
				item *= m.Int(monkey.Operation[2])
			default:
				log.Fatalln("unknown operation", monkey.Operation)
			}

			if scale == 1 {
				// special handling if scale is one, reduce by gcf
				item = item % gcf
			} else {
				item /= scale
			}

			// check if it's divisible by the test
			target := monkey.False
			if item%monkey.Test == 0 {
				target = monkey.True
			}
			k := monkeys[target]

			k.Items = append(k.Items, item)
		}
		monkey.Items = []int{} // clear the monkey's items
	}

}

type Monkey struct {
	Items     []int
	Operation []string
	Test      int // always divisible???
	True      int
	False     int

	Count int
}

func loadInput(s string) []*Monkey {
	if s != "" {
		input.SetRaw(s)
	}
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	things := input.LoadSliceString("\n\n")

	var list []*Monkey
	for _, ms := range things {
		lines := strings.Split(ms, "\n")
		x := strings.TrimPrefix(lines[1], "  Starting items: ")
		items := m.IntSlice(strings.Split(x, ", "))
		operation := strings.Split(strings.TrimPrefix(lines[2], "  Operation: new = "), " ")
		test := m.Int(strings.TrimPrefix(lines[3], "  Test: divisible by "))
		t := m.Int(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
		f := m.Int(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

		m := Monkey{
			Items:     items,
			Operation: operation,
			Test:      test,
			True:      t,
			False:     f,
		}
		list = append(list, &m)
	}
	return list
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
