#!/bin/env python3

import sys
import re


def main(): 
    if len(sys.argv) != 2:
        print(f"[usage]: {__file__} [puzzle]", file=sys.stderr)
        sys.exit(1)

    max = (12, 13, 14)

    nopos_l = []
    total_game = 0

    with open(sys.argv[1]) as f: 
        for i, line in enumerate(f): 
            game = re.search(r"Game (\d+): ", line)
            if game == None: 
                print(f"[ERRO]: no game information on line {i:{3}}", file=sys.stderr)
                sys.exit(1)
            game_num = int(game.group(1))
            total_game = game_num
            line = line[game.span()[1]:]

            # get all the numbers
            strs = line.split(";")
            for s in strs: 
                d = [0, 0, 0]

                r = re.findall(r"(\d+) red", s)
                g = re.findall(r"(\d+) green", s)
                b = re.findall(r"(\d+) blue", s)
                
                # fill in table
                for t in r:
                    d[0] += int(t)
                for t in g: 
                    d[1] += int(t)
                for t in b: 
                    d[2] += int(t)
                
                # check if in bouds
                if not(d[0] <= max[0] and d[1] <= max[1] and d[2] <= max[2]): 
                    nopos_l.append(game_num)
                    break


    # inverse `nopos_l` to get all the possible ones
    pos_l = list(range(1, total_game+1))
    pos_l = [x for x in pos_l if x not in nopos_l]
    # print(pos_l)

    sum = 0
    for x in pos_l: 
        sum += x
    print(sum)




if __name__ == "__main__": 
    main()
