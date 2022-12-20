package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "unicode"
    "strconv"
    "strings"
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
    
    procedure := false
    stacks := make([][]rune, 0)

    var s string
    for scanner.Scan() {
        s = scanner.Text()

        if len(s) == 0 {
            procedure = true
            continue
        }

        if len(stacks) == 0 {
            n := (len(s) + 1) / 4
            for i := 0; i < n; i++ {
                temp := make([]rune, 0)
                stacks = append(stacks, [][]rune{temp}...)
            }
        }

        // starting stacks of crates
        if !procedure {
            chars := []rune(s)
            for i := 1; i < len(chars); i += 4 {
                if !unicode.IsSpace(chars[i]) && !unicode.IsNumber(chars[i]) {
                    stackNum := (i - 1) / 4

                    crate := chars[i]
                    stacks[stackNum] = append([]rune{crate}, stacks[stackNum]...)
                }
            }

            continue
        }

        // rearrangement procedure
        instruction := strings.Fields(s)

        num, err := strconv.Atoi(instruction[1])
        check(err)

        startStack, err := strconv.Atoi(instruction[3])
        check(err)
    
        endStack, err := strconv.Atoi(instruction[5])
        check(err)
        
        startStack -= 1
        endStack -= 1

        // Part 1
        if *partPtr == 1 {
            for i := 0; i < num && len(stacks[startStack]) > 0; i++ {
                crate := stacks[startStack][len(stacks[startStack]) - 1]
                
                stacks[startStack] = stacks[startStack][:len(stacks[startStack])-1]
                stacks[endStack] = append(stacks[endStack], []rune{crate}...)
            }

            continue
        }

        // Part 2
        stacks[endStack] = append(stacks[endStack], stacks[startStack][len(stacks[startStack])-num:]...)
        stacks[startStack] = stacks[startStack][:len(stacks[startStack])-num]
        
    }

    err = scanner.Err()
    check(err)

    var topCrates []rune
    for i := 0; i < len(stacks); i++ {
        topCrates = append(topCrates, stacks[i][len(stacks[i])-1])
    }

    answer := string(topCrates)

    fmt.Printf("Input File: %s\n", *fNamePtr)
    fmt.Printf("Part %d: %s\n", *partPtr, answer)
}
