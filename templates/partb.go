package main

import (
	"fmt"

	"github.com/jnewmano/advent2022/input"
	m "github.com/jnewmano/advent2022/simplemath"
)

func b() {
	s := ``
	s = ``

	things := loadInput(s)
	list := process(things)

	fmt.Println("sum of list :", m.Sum(list))
	fmt.Println("high in list:", m.High(list))
	fmt.Println("low in list :", m.Low(list))
	fmt.Println("len of list :", len(list))

}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
