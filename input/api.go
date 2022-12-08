package input

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func loadInputFromAPI() {
	path, _ := os.Getwd()
	inputFilePath := filepath.Join(path, "input.txt")

	day := filepath.Base(path)
	if !strings.HasPrefix(day, "day") {
		log.Fatalln("invalid path, expected day in direction name")
	}

	day = strings.TrimPrefix(day, "day")
	year := "2022"

	fmt.Printf("loading remote input file for %s day %s\n", year, day)

	cmd := exec.Command("aocdl", "-year", year, "-day", day, "-output", inputFilePath)
	fmt.Println("%", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("Unable to load input data:", strings.TrimSpace(string(output)))
	}

	fmt.Println(string(output))
}

func checkSHA(output []byte) bool {

	expectedSHA := sha256.Sum256(output)

	_, err := os.Stat("input.txt.sha")
	if os.IsNotExist(err) {
		writeSHA(expectedSHA[:])
	}

	sha, err := os.ReadFile("input.txt.sha")
	if err != nil {
		panic(err)
	}

	if !bytes.Equal(sha, expectedSHA[:]) {
		fmt.Printf("got:      %x\n", sha)
		fmt.Printf("expected: %x\n", expectedSHA)
		log.Fatal("input.txt sha is different than expected")
	}

	return true
}

func writeSHA(sha []byte) {
	err := os.WriteFile("input.txt.sha", sha, 0444)
	if err != nil {
		log.Fatal(err)
	}
}
