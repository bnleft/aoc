package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func priority(item byte) int {
    var res int
    res = int(item)

    if item >= byte('a') && item <= byte('z') {
        res += 1 - int('a')
    } else if item >= byte('A') && item <= byte('Z') {
        res += 27 - int('A')
    }
    
    return res
}

func main() {
    fName := string(os.Args[1])

    f, err := os.Open(fName)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)

    total := 0

    var sack string
    for scanner.Scan() {
        sack = scanner.Text()
        half := len(sack) / 2
        comp1 := []rune(sack[:half])
        comp2 := []rune(sack[half:])

        var common string
        for i := 0; i < len(comp1); i++ {
            for j := 0; j < len(comp2); j++ {
                if comp1[i] == comp2[j] {
                    common = string(comp1[i])
                }
            }
        }

        total += priority(common[0])
    }

    err = scanner.Err()
    check(err)

    fmt.Printf("Part 1: %d\n", total)
}
