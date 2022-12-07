package input

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2021/simplemath"
)

var raw string

// SetRaw lets you set the raw input, overriding the AoC provided input
func SetRaw(a string) error {
	raw = strings.TrimSpace(a)
	return nil
}

// Load returns the input file as a string
func Load() string {
	fn := "input.txt"
	useExample, _ := strconv.ParseBool(os.Getenv("E"))
	if useExample {
		fn = "example.txt"
	}

	if raw != "" {
		return raw
	}

	_, err := os.Stat(fn)
	if os.IsNotExist(err) {
		loadInputFromAPI()
	}

	all, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	if !useExample && !checkSHA(all) {
		panic("sha didn't validate, input.txt file has been modified")
	}

	return strings.TrimSpace(string(all))
}

// LoadSliceString returns a slice of strings
// defaults to newline deliminated
func LoadSliceString(sep string) []string {
	if sep == "" {
		sep = "\n"
	}
	all := Load()
	things := strings.Split(all, sep)
	return things
}

// LoadSliceString returns a slice of ints
// defaults to newline deliminated
func LoadSliceInt(sep string) []int {
	things := LoadSliceString(sep)
	return simplemath.IntSlice(things)
}

// LoadSliceSliceString does a double split on sepLine and then on sepRow
func LoadSliceSliceString(sepLine string, sepRow string) [][]string {
	if sepRow == "" {
		sepRow = ","
	}

	things := LoadSliceString(sepLine)
	var resp [][]string

	for _, v := range things {
		v = strings.TrimSpace(v)
		parts := strings.Split(v, sepRow)
		resp = append(resp, parts)
	}
	return resp
}
