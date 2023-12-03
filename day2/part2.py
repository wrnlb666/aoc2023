#!/bin/env python3

import sys
import re


def main(): 
    if len(sys.argv) != 2:
        print(f"[usage]: {__file__} [puzzle]", file=sys.stderr)
        sys.exit(1)


    sum = 0

    with open(sys.argv[1]) as f: 
        for i, line in enumerate(f): 
            game = re.search(r"Game (\d+): ", line)
            if game == None: 
                print(f"[ERRO]: no game information on line {i:{3}}", file=sys.stderr)
                sys.exit(1)
            game_num = int(game.group(1))
            line = line[game.span()[1]:]

            # get all the numbers
            d = [0, 0, 0]
            strs = line.split(";")
            for s in strs: 
                temp = [0, 0, 0]

                r = re.findall(r"(\d+) red", s)
                g = re.findall(r"(\d+) green", s)
                b = re.findall(r"(\d+) blue", s)
                
                # fill in table
                for t in r:
                    temp[0] += int(t)
                for t in g: 
                    temp[1] += int(t)
                for t in b: 
                    temp[2] += int(t)
                
                d[0] = d[0] if temp[0] <= d[0] else temp[0]
                d[1] = d[1] if temp[1] <= d[1] else temp[1]
                d[2] = d[2] if temp[2] <= d[2] else temp[2]                
        
            sum += d[0] * d[1] * d[2]

        print(sum)




if __name__ == "__main__": 
    main()

