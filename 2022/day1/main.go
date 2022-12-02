package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
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

    curNum := 0
    mostCal := 0

    var s string
    for scanner.Scan() {
        s = scanner.Text()
        if len(s) == 0 {
            curNum = 0
            continue
        }
        
        cal, err := strconv.Atoi(s)
        check(err)

        curNum += cal
        if curNum > mostCal {
            mostCal = curNum
        }
    }

    err = scanner.Err()
    check(err)

    fmt.Printf("Most Calories: %d", mostCal)
}
