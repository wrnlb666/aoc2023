#define _GNU_SOURCE
#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdbool.h>


int main(int argc, char* argv[]) {
    if (argc != 2) {
        fprintf(stderr, "[usage]: %s [puzzle]\n", argv[0]);
        exit(1);
    }

    FILE* fp = fopen(argv[1], "r");
    if (fp == NULL) {
        fprintf(stderr, "[ERRO]: failed to open file %s\n", argv[1]);
        exit(1);
    }

    int line = 0;
    int left;
    int right;

    int32_t curr;
    int32_t sum = 0;

    size_t buf_size = 100;
    char* buf = malloc(sizeof (char) * buf_size);

    while (true) {
        line += 1;
        // get the current line
        int len = getline(&buf, &buf_size, fp);
        // buf = fgets(buf, buf_size, fp);

        // break if eof
        if (len == -1) {
            break;
        }
        // if (buf == NULL) {
        //     break;
        // }
        // int len = strlen(buf);

        // set left and right to initial state
        left = -1;
        right = -1;
        
        for (int i = 0; i < len; i++) {
            if (isdigit(buf[i])) {
                if (left == -1) {
                    left = buf[i] - 0x30;
                } else {
                    right = buf[i] - 0x30;
                }
            }        
        }
        if (left == -1) {
            fprintf(stderr, "[ERRO]: no number at line %d\n", line);
            exit(1);
        }
        // if only one digit
        if (right == -1) {
            right = left;
        }
        
        // calculate result
        curr = left * 10 + right;
        // printf("line %3d, curr: %4d, str: %.*s\n", line, curr, len - 1, buf);
        sum += curr;
    }
    free(buf);
    fclose(fp);
    printf("sum: %d\n", sum);

    return 0;
}

