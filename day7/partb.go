package main

import (
	"fmt"
	"sort"

	"github.com/jnewmano/advent2021/input"
	m "github.com/jnewmano/advent2021/simplemath"
)

func b() {
	s := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	s = ``

	things := loadInput(s)
	list := process(things)

	used := sizes(&list)

	totalDisk := 70000000
	needed := 30000000

	toFree := needed - (totalDisk - used)

	fmt.Println("To free:", toFree)

	sort.Ints(dirSizes)
	for _, v := range dirSizes {
		if v >= toFree {
			fmt.Println("Smallest dir:", v)
			break
		}
		fmt.Println("current:", v)
	}

}

var (
	_ = input.LoadSliceString
	_ = m.Abs
)
