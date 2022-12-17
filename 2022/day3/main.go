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
    total2 := 0

    var sack string
    var group []string
    for scanner.Scan() {
        sack = scanner.Text()

        // Part 1
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

        // Part 2
        group = append(group, sack)

        if len(group) == 3 {
            elf1 := []rune(group[0])
            elf2 := []rune(group[1])
            elf3 := []rune(group[2])

            var common2 string
            for i := 0; i < len(elf1); i++ {
                for j := 0; j < len(elf2); j++ {
                    for k := 0; k < len(elf3); k++ {
                        if elf1[i] == elf2[j] && elf1[i] == elf3[k] {
                            common2 = string(elf1[i])
                        }
                    }
                }
            }

            total2 += priority(common2[0])
            
            group = nil
        }
    }

    err = scanner.Err()
    check(err)

    fmt.Printf("Part 1: %d\n", total)
    fmt.Printf("Part 2: %d\n", total2)
}
