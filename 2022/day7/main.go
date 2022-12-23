package main

import (
	"bufio"
	"flag"
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

type Dir struct {
    parent *Dir
    subs []*Dir
    name string
    size int
}

func adjustSize(parent *Dir) {
    if parent == nil {
        return
    }

    for _, subDir := range parent.subs {
        adjustSize(subDir)
    }

    for _, subDir := range parent.subs {
        parent.size += subDir.size
    }
}

func sumSizeMost(parent *Dir, answer *int, most int) {
    if parent == nil {
        return
    }

    for _, subDir := range parent.subs {
        sumSizeMost(subDir, answer, most)
    }

    if parent.size <= most {
        *answer += parent.size
    }
}

func smallestUnused(parent *Dir, answer *int, smallest int) {
    if parent == nil {
        return
    }

    for _, subDir := range parent.subs {
        smallestUnused(subDir, answer, smallest)
    }

    if parent.size >= smallest && parent.size < *answer {
        *answer = parent.size
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
    
    answer := 0

    root := &Dir{
        name: "/",
        size: 0,
    }
    curDir := root

	var s string
	for scanner.Scan() {
		s = scanner.Text()
        line := strings.Fields(s)
    
        // Command
        if line[0] == "$" {

            // Change Directory
            if line[1] == "cd" {
                newDir := line[2]

                // Move back one level
                if newDir == ".." {
                    curDir = curDir.parent
                    continue
                }

                // Switch to outermost directory
                if newDir == "/" {
                    curDir = root
                    continue
                }

                // Move down one level 
                for _, subDir := range curDir.subs {
                    if subDir.name == newDir {
                        curDir = subDir
                        break
                    }
                }

                continue
            }
            
            // List contents
            if line[1] == "ls" {
                continue
            }
        }

        // Directory
        if line[0] == "dir" {
            subDirName := line[1]
            // Add subdirectory to current directory
            subDir := &Dir{
                parent: curDir,
                name: subDirName,
                size: 0,
            }
            curDir.subs = append(curDir.subs, subDir)

            continue
        }

        // File
        fileSize, err := strconv.Atoi(line[0])
        check(err)

        curDir.size += fileSize
	}

	err = scanner.Err()
	check(err)

    // Adjust Directories Sizes
    adjustSize(root)

    // Part 1
    if *partPtr == 1 {
        most := 100000
        sumSizeMost(root, &answer, most)
    }

    // Part 2
    if *partPtr == 2 {
        totalDisk := 70000000
        needed := 30000000 
        unused := totalDisk - root.size
        smallest := needed - unused
        answer = needed

        smallestUnused(root, &answer, smallest)
    }

	fmt.Printf("Input File: %s\n", *fNamePtr)
	fmt.Printf("Part %d: %d\n", *partPtr, answer)
}
