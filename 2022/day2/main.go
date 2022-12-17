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

func shape(s string) int {
    val := -1

    if s == "A" || s == "X" { // Rock
        val = 1
    } else if s == "B" || s == "Y" { // Paper
        val = 2
    } else if s == "C" || s == "Z" { // Scissors
        val = 3
    }

    return val
}

func outcome(p1 int, me int) int {
    // Lose
    val := 0

    if p1 == me { // Draw
        val = 3
    } else if (p1 == 1 && me == 2) || (p1 == 2 && me == 3) || (p1 == 3 && me == 1) { // Win
        val = 6
    } 

    return val
}

func choose(op int, end string) int {
    res := 0

    if end == "X" { // Lose
        if op == 1 {
            res += 3
        } else if op == 2 {
            res += 1
        } else if op == 3 {
            res += 2
        }
    } else if end == "Y" { // Draw
        res += op + 3
    } else if end == "Z" { // Win
        if op == 1 {
            res += 2
        } else if op == 2 {
            res += 3
        } else if op == 3 {
            res += 1
        }

        res += 6
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
    total2 := 0

    var s string
    for scanner.Scan() {
        s = scanner.Text()
        p1 := string(s[0])
        p2 := string(s[2])

        total += shape(p2) + outcome(shape(p1), shape(p2))

        total2 += choose(shape(p1), p2) 
    }

    err = scanner.Err()
    check(err)

    fmt.Printf("Part 1: %d\n", total)
    fmt.Printf("Part 2: %d\n", total2)
}
