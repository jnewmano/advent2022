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
	s := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	s = ``

	things := loadInput(s)
	n := process(4, things)
	fmt.Println(n)

}

func loadInput(s string) string {
	if s != "" {
		input.SetRaw(s)
	}

	return input.Load()
}

func process(n int, data string) int {

	for i := range data {
		if i < n {
			continue
		}
		m := map[byte]struct{}{}
		for j := 0; j < n; j++ {
			k := i - j
			m[data[k]] = struct{}{}
		}
		if len(m) == n {
			return (i + 1)
		}
	}

	return 0
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

var (
	_ = m.Abs
)
