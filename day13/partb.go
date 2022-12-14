package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func b() {
	s := ``
	s = ``

	things := flatten(loadInput(s))
	packetA := "[[2]]"
	packetB := "[[6]]"

	things = append(things, packetA)
	things = append(things, packetB)

	sort.Sort(packetSort(things))

	var idxA, idxB int
	for i, v := range things {
		if v == packetA {
			idxA = i + 1
		}
		if v == packetB {
			idxB = i + 1
		}
	}
	fmt.Println(idxA, idxB, idxA*idxB)
}

type packetSort []string

func (s packetSort) Less(i, j int) bool {
	var a, b any
	json.Unmarshal([]byte(s[i]), &a)
	json.Unmarshal([]byte(s[j]), &b)

	c := compare(a, b)
	return c == 1
}

func (s packetSort) Len() int {
	return len(s)
}

func (s packetSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func flatten(list [][]string) []string {
	flat := []string{}
	for _, v := range list {
		flat = append(flat, v...)
	}
	return flat
}
