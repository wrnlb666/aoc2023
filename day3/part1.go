package main

import (
    "os"
    "io"
    "fmt"
    "unicode"
    "strings"
    "strconv"
)

func part1() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "[usage]: %s [puzzle]\n", os.Args[0]);
        os.Exit(1)
    }
    
    // open file
    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERRO]: %s\n", err)
        os.Exit(1)
    }
    defer f.Close()

    // read file
    file, err := io.ReadAll(f)
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERRO]: %s\n", err)
        os.Exit(1)
    }

    // convert to matrix
    var m matrix
    m.data = strings.Split(string(file), "\n")
    if len(m.data[len(m.data)-1]) == 0 {
        m.data = m.data[:len(m.data)-1]
    }
    m.height = len(m.data) - 1
    m.width = len(m.data[0]) - 1

    var s []int
    curr := -1
    added := false

    // parsing
    for x := 0; x <= m.height; x++ {
        for y := 0; y <= m.width; y++ {
            r := rune(m.data[x][y])
            if unicode.IsDigit(r) {
                if curr == -1 {
                    currStr := m.data[x][y:]
                    numStr := ""
                    for _, char := range currStr {
                        if unicode.IsDigit(char) {
                            numStr += string(char)
                        } else {
                            break
                        }
                    }
                    curr, err = strconv.Atoi(numStr)
                    if err != nil {
                        fmt.Println(err)
                    }
                }
                if m.p1Adjcent(x, y) {
                    if added == false {
                        s = append(s, curr)
                        added = true
                    }
                }
            } else {
                curr = -1
                added = false
            }
        }
    }

    sum := 0
    for _, v := range s {
        // fmt.Println(v)
        sum += v
    }
    // fmt.Printf("len: %d\n", len(s))
    fmt.Printf("part1: %d\n", sum)
}

func p1Check(r rune) bool {
    if r != '.' && !unicode.IsDigit(r) {
        return true
    }
    return false
}

func (m matrix) p1Adjcent(x, y int) bool {
    // check all 8 position around (x, y)
    var curr rune
    if x != 0 {
        curr = rune(m.data[x-1][y])
        if p1Check(curr) {
            return true
        }
        if y != 0 {
            curr = rune(m.data[x-1][y-1])
            if p1Check(curr) {
                return true
            }
        }
        if y != m.width {
            curr = rune(m.data[x-1][y+1])
            if p1Check(curr) {
                return true
            }
        }
    }

    if x != m.height {
        curr = rune(m.data[x+1][y])
        if p1Check(curr) {
            return true
        }
        if y != 0 {
            curr = rune(m.data[x+1][y-1])
            if p1Check(curr) {
                return true
            }
        }
        if y != m.width {
            curr = rune(m.data[x+1][y+1])
            if p1Check(curr) {
                return true
            }
        }
    }

    if y != 0 {
        curr = rune(m.data[x][y-1])
        if p1Check(curr) {
            return true
        }
    }
    if y != m.width {
        curr = rune(m.data[x][y+1])
        if p1Check(curr) {
            return true
        }
    }

    return false
}
