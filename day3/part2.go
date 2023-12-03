package main

import (
    "os"
    "io"
    "fmt"
    "unicode"
    "strings"
    "strconv"
)

type pos struct {
    x   int
    y   int
}

func part2() {
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

    // var g []gear
    g := make(map[pos][]int)
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
                gx, gy := m.p2Adjcent(x, y)
                if gx != -1 && gy != -1 {
                    if added == false {
                        g[pos{gx, gy}] = append(g[pos{gx, gy}], curr)
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
    for _, v := range g {
        // assume at most 2 part numbers will be adjcent to the same gear
        if len(v) == 2 {
            sum += v[0] * v[1]
        }
    }
    // fmt.Printf("len: %d\n", len(s))
    fmt.Printf("part2: %d\n", sum)
}

func p2Check(r rune) bool {
    if r == '*' {
        return true
    }
    return false
}

// return the (x,y) pos of that `*`
func (m matrix) p2Adjcent(x, y int) (int, int) {
    // check all 8 position around (x, y)
    var curr rune
    if x != 0 {
        curr = rune(m.data[x-1][y])
        if p2Check(curr) {
            return x-1, y
        }
        if y != 0 {
            curr = rune(m.data[x-1][y-1])
            if p2Check(curr) {
                return x-1, y-1
            }
        }
        if y != m.width {
            curr = rune(m.data[x-1][y+1])
            if p2Check(curr) {
                return x-1, y+1
            }
        }
    }

    if x != m.height {
        curr = rune(m.data[x+1][y])
        if p2Check(curr) {
            return x+1, y
        }
        if y != 0 {
            curr = rune(m.data[x+1][y-1])
            if p2Check(curr) {
                return x+1, y-1
            }
        }
        if y != m.width {
            curr = rune(m.data[x+1][y+1])
            if p2Check(curr) {
                return x+1, y+1
            }
        }
    }

    if y != 0 {
        curr = rune(m.data[x][y-1])
        if p2Check(curr) {
            return x, y-1
        }
    }
    if y != m.width {
        curr = rune(m.data[x][y+1])
        if p2Check(curr) {
            return x, y+1
        }
    }

    return -1, -1
}
