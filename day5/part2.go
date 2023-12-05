package main

import (
    "os"
    "io"
    "log"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "[usage]: %s [puzzle]\n", os.Args[0])
        os.Exit(1)
    }

    var seeds []int64
    var maps [][][3]int64

    // parse file
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    f, err := io.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    lines := strings.Split(string(f), "\n")

    state := -1
    started := false

    for _, line := range lines {
        if state == -1 {
            if len(line) == 0 {
                state += 1
            } else {
                s := strings.Split(line, " ")[1:]
                for _, n := range s {
                    res, err := strconv.ParseInt(n, 10, 64)
                    if err != nil {
                        log.Fatal(err)
                    }
                    seeds = append(seeds, res)
                }
            }
        } else {
            if started == false {
                started = true
                maps = append(maps, [][3]int64{})
                continue
            } else {
                if len(line) == 0 {
                    state += 1
                    started = false
                } else {
                    s := strings.Split(line, " ")
                    n1, err := strconv.ParseInt(s[0], 10, 64)
                    if err != nil {
                        log.Fatal(err)
                    }
                    n2, err := strconv.ParseInt(s[1], 10, 64)
                    if err != nil {
                        log.Fatal(err)
                    }
                    n3, err := strconv.ParseInt(s[2], 10, 64)
                    if err != nil {
                        log.Fatal(err)
                    }
                    maps[state] = append(maps[state], [3]int64{n1, n2, n3})
                }
            }
        }
    }

    var min_num int64 = 9999999999
    var res [][2]int64
    for i := 0; i < len(seeds); i+= 2 {
        for s := seeds[i]; s < seeds[i] + seeds[i+1]; s++ {
            curr := s
            for i := range res {
                if s >= res[i][0] && s < res[i][1] {
                    continue
                }
            }
            for i := range maps {
                for j := range maps[i] {
                    if curr >= maps[i][j][1] && curr < maps[i][j][1] + maps[i][j][2] {
                        offset := curr - maps[i][j][1]
                        curr = maps[i][j][0] + offset
                        break
                    }
                }
            }
            if curr < min_num {
                min_num = curr
            }
        }
        res = append(res, [2]int64{seeds[i], seeds[i]+seeds[i+1]})
        fmt.Println("done")
    }

    fmt.Println(min_num)
}

