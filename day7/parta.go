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

	_ = sizes(&list)

	fmt.Println(grandTotal)
}

var grandTotal int
var dirSizes []int

func sizes(dir *Dir) int {
	var total int
	for _, v := range dir.Things {
		switch tt := v.(type) {
		case *Dir:
			total += sizes(tt)
		case *File:
			total += tt.Size
		default:
			log.Fatalf("unhandled type of thing [%T]\n", dir.Things)
		}
	}

	if total <= 100000 {
		grandTotal += total
	}
	dirSizes = append(dirSizes, total)

	return total
}

func loadInput(s string) []string {
	if s != "" {
		input.SetRaw(s)
	}
	return input.LoadSliceString("")
}

type Dir struct {
	Name   string
	Things map[string]interface{}
	Parent *Dir
}

type File struct {
	Name string
	Size int
}

func process(things []string) Dir {

	root := &Dir{
		Name:   "root",
		Things: make(map[string]interface{}),
		Parent: nil,
	}

	current := root

	for _, v := range things {
		if strings.HasPrefix(v, "$ cd") {
			// check to see if we have a directory for it
			target := strings.TrimPrefix(v, "$ cd ")
			switch target {
			case "/":
				current = root
			case "..":
				current = current.Parent
			default:
				if _, ok := current.Things[target]; ok {
					current = current.Things[target].(*Dir)
				} else {
					new := Dir{
						Name:   target,
						Things: make(map[string]interface{}),
						Parent: current,
					}
					current.Things[target] = new
					current = &new
				}
			}

		} else if strings.HasPrefix(v, "$ ls") {
		} else {
			// assume we're listing files
			if strings.HasPrefix(v, "dir") {
				dn := strings.TrimPrefix(v, "dir ")
				new := &Dir{
					Name:   dn,
					Things: make(map[string]interface{}),
					Parent: current,
				}
				current.Things[dn] = new
			} else {
				size, fn, _ := strings.Cut(v, " ")
				file := &File{
					Size: m.Int(size),
					Name: fn,
				}
				current.Things[fn] = file
			}
		}
	}

	return *root
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
