package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fNamePtr := flag.String("file", "example.txt", "a string")
	partPtr := flag.Int("part", 1, "an int")

	flag.Parse()

	f, err := os.Open(*fNamePtr)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var answer string

	var s string
	for scanner.Scan() {
		s = scanner.Text()
        answer = s
	}

	err = scanner.Err()
	check(err)

	fmt.Printf("Input File: %s\n", *fNamePtr)
	fmt.Printf("Part %d: %s\n", *partPtr, answer)
}
