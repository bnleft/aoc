package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fName := string(os.Args[1])

    f, err := os.Open(fName)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)

    total := 0

    var pair string
    for scanner.Scan() {
        pair = scanner.Text()

        elves := strings.Split(pair, ",")
        
        sects1 := strings.Split(elves[0], "-")
        start1, err := strconv.Atoi(sects1[0])
        check(err)
        end1, err := strconv.Atoi(sects1[1])
        check(err)

        sects2 := strings.Split(elves[1], "-")
        start2,err := strconv.Atoi(sects2[0])
        check(err)
        end2, err := strconv.Atoi(sects2[1])
        check(err)

        if start1 >= start2 && end1 <= end2 { // sects1 within
            total += 1
        } else if start2 >= start1 && end2 <= end1 { // sects2 within
            total += 1
        }
        
    }

    err = scanner.Err()
    check(err)
        
    fmt.Printf("Part 1: %d\n", total)
}
