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

    loc = []
    for s in seeds: 
        curr = s
        for m in maps: 
            for l in m: 
                if curr >= l[1] and curr < l[1] + l[2]: 
                    offset = curr - l[1]
                    curr = l[0] + offset
                    break
        loc.append(curr)

    print(min(loc))



if __name__ == "__main__": 
    main()
