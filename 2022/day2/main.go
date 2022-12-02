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

func main() {
    fName := string(os.Args[1])

    f, err := os.Open(fName)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)

    total := 0

    var s string
    for scanner.Scan() {
        s = scanner.Text()
        p1 := string(s[0])
        p2 := string(s[2])

        total += shape(p2) + outcome(shape(p1), shape(p2))
    }

    err = scanner.Err()
    check(err)


    fmt.Printf("%d\n", total)
}
