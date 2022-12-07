package main

import (
	"fmt"
	"strings"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func b() {
	s := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	s = ``

	things := loadInput(s)
	list := processB(things)

	// fmt.Println("sum of list :", m.Sum(list))
	// fmt.Println("high in list:", m.High(list))
	// fmt.Println("low in list :", m.Low(list))
	fmt.Println("len of list :", len(list))
}

func processB(things [][]string) []int {

	list := []int{}
	for i, v := range things {
		a := strings.Split(v[0], "-")
		b := strings.Split(v[1], "-")
		a1 := m.Int(a[0])
		a2 := m.Int(a[1])

		b1 := m.Int(b[0])
		b2 := m.Int(b[1])

		if a1 <= b1 {
			if a2 >= b1 {
				fmt.Println("a", i)
				list = append(list, i)
			}
			continue
		}

		if b1 < a1 {
			if b2 >= a1 {
				fmt.Println("b", i)
				list = append(list, i)
			}
			continue
		}
	}

	/*
		2-4,6-8 no
		2-3,4-5 no
		5-7,7-9 yes
		2-8,3-7 yes
		6-6,4-6 yes
		2-6,4-8 yes

	*/
	return list
}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
