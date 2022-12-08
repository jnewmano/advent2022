package main

import (
	"fmt"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func b() {
	s := ``

	things := loadInput(s)
	list := processB(things)

	max := 0
	for _, v := range list {
		for _, vv := range v {
			if vv > max {
				max = vv
			}
		}
	}
	fmt.Println(max)
}

func processB(things []string) [][]int {
	scores := m.New2DSlice[int](len(things), len(things[0]))

	for i := 0; i < len(things); i++ {
		for j := 0; j < len(things); j++ {

			// calculate the distance that can be seen from this point
			c := h(things, i, j)

			// check moving left
			left := 0
			for y := j - 1; y >= 0; y-- {
				nx := h(things, i, y)
				left++
				if nx >= c {
					break
				}
			}

			// check moving right
			right := 0
			for y := j + 1; y < len(things); y++ {
				nx := h(things, i, y)
				right++
				if nx >= c {
					break
				}
			}

			// check moving up
			up := 0
			for x := i - 1; x >= 0; x-- {
				nx := h(things, x, j)
				up++
				if nx >= c {
					break
				}
			}

			// check moving down
			down := 0
			for x := i + 1; x < len(things); x++ {
				nx := h(things, x, j)
				down++
				if nx >= c {
					break
				}
			}

			score := m.Product(left, right, up, down)
			scores[i][j] = score
		}
	}
	return scores
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
