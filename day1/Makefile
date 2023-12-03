CC = gcc
CFLAG = -Wall -Wextra -std=c11 -pedantic -g

.PHONY: test

test: test.c 
	$(CC) $(CFLAG) $< -o $@
	./$@ puzzle.txt
