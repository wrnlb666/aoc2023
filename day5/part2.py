#!/usr/bin/env python3

import sys

def main(): 
    if len(sys.argv) != 2: 
        print(f"[usage]: {sys.argv[0]} [puzzle.txt]", file=sys.stderr)

    seeds = []
    maps = []

    # parse file
    with open(sys.argv[1], "r") as f: 
        state = -1
        started = False

        for line in f: 
            if state == -1: 
                if len(line) == 1: 
                    state += 1
                s = line.split(" ")[1:]
                for n in s: 
                    seeds.append(int(n))

            else: 
                if started == False: 
                    started = True 
                    maps.append([])
                    continue 
                elif started == True: 
                    if len(line) == 1: 
                        state += 1
                        started = False 
                    else: 
                        s = line.split(" ")
                        res = list(map(int, s))
                        maps[state].append(res)


    min = 999999999
    res = []
    for i in range(0, len(seeds), 2): 
        for s in range(seeds[i], seeds[i]+seeds[i+1]): 
            curr = s
            for x in res: 
                if s >= x[0] and s < x[1]: 
                    continue
            for m in maps: 
                for l in m: 
                    if curr >= l[1] and curr < l[1] + l[2]: 
                        offset = curr - l[1]
                        curr = l[0] + offset
                        break
            if curr < min: 
                min = curr
        res.append((seeds[i], seeds[i]+seeds[i+1]))
        print("done")

    print(min)



if __name__ == "__main__": 
    main()

