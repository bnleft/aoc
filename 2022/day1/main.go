package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
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

    
    var s string
    var foods []int
    curFoodCal := 0

    for scanner.Scan() {
        s = scanner.Text()
        if len(s) == 0 {
            foods = append(foods, curFoodCal)
            curFoodCal = 0
            continue
        }
        
        cal, err := strconv.Atoi(s)
        check(err)

        curFoodCal += cal
    }

    err = scanner.Err()
    check(err)

    sort.Ints(foods)

    last := len(foods) - 1

    mostCal := foods[last]
    sumCal := 0

    for i := last; i >= 0 && i > last - 3; i-- {
        sumCal += foods[i]
    }

    fmt.Printf("Most Calories: %d\n", mostCal)
    fmt.Printf("Total Calories of Top 3 Food: %d\n", sumCal)
}
