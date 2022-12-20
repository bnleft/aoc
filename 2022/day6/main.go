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

	var answer int
    distinct := 4

    if *partPtr != 1 {
        distinct = 14
    }

	for scanner.Scan() {
		buffer := []rune(scanner.Text())

        var marker int

        chars := map[rune]int{}
        count := 0
        start := 0
        for marker = 0; marker < len(buffer); marker++ {
            val, ok := chars[buffer[marker]]

            chars[buffer[marker]] = marker
            count += 1

            if ok && val >= start {
                start = val + 1
                count = marker - start + 1
                continue
            }
            
            
            if count == distinct {
                break
            }
        }
        
        answer = marker + 1
	}

	err = scanner.Err()
	check(err)

	fmt.Printf("Input File: %s\n", *fNamePtr)
	fmt.Printf("Part %d: %d\n", *partPtr, answer)
}
