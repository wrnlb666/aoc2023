CC = gcc
CFLAG = -Wall -Wextra -std=c11 -pedantic -g

.PHONY: part1 part2

part1: part1.c 
	$(CC) $(CFLAG) $< -o $@
	./$@ puzzle.txt

part2: part2.c 
	$(CC) $(CFLAG) $< -o $@
	./$@ puzzle.txt

