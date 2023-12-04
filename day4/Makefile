CC = gcc
CFLAG = -Wall -Wextra -std=c11 -pedantic -g
LIB = -lm

.PHONY: part1 part2

part1: part1.c 
	$(CC) $(CFLAG) $< -o $@ $(LIB)
	./$@ puzzle.txt

part2: part2.c 
	$(CC) $(CFLAG) $< -o $@ $(LIB) #-fopenmp
	./$@ puzzle.txt

