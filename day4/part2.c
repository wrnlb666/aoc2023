#define _GNU_SOURCE
#include <math.h>
#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdbool.h>


typedef struct slice {
    int*    arr;
    int     len;
    int     cap;
} slice_t;


void append(slice_t* s, int c) {
    if (s->arr == NULL) {
        s->cap = 16;
        s->arr = malloc(sizeof (int) * 16);
    } else if (s->len + 1 == s->cap) {
        s->cap *= 2;
        s->arr = realloc(s->arr, sizeof (int) * s->cap);
    }
    s->arr[s->len] = c;
    s->len += 1;
}


int check_card(const slice_t* s, int index) {
    int sum = 1;
    if (index + 1 == s->len) {
        return sum;
    }
    for (int i = index + 1; i < index + s->arr[index] + 1 && i < s->len; i++) {
        sum += check_card(s, i);
    }
    return sum;
}


int main(int argc, char** argv) {
    if (argc != 2) {
        fprintf(stderr, "[usage]: %s [puzzle]\n", argv[0]);
        exit(1);
    }

    FILE* fp = fopen(argv[1], "r");

    size_t buf_size = 50;
    char* buf = malloc(sizeof (char) * buf_size);

    // parse first line to check size
    int win_num = 0;
    int my_num = 0;
    bool is_num = false;
    int len = getline(&buf, &buf_size, fp);
    int state = 0;
    for (int i = 0; i < len; i++) {
        char c = buf[i];
        switch (state) {
            case 0: {
                if (c == ':') {
                    state = 1;
                }
                break;
            }
            case 1: {
                if (c == '|') {
                    state = 2;
                } else if (isdigit(c)) {
                    if (is_num == false) {
                        is_num = true;
                        win_num += 1;
                    }
                } else {
                    is_num = false;
                }
                break;
            }
            case 2: {
                if(isdigit(c)) {
                    if (is_num == false) {
                        is_num = true;
                        my_num += 1;
                    }
                } else {
                    is_num = false;
                }
                break;
            }
        }
    }

    int* win_nums = malloc(sizeof (int) * win_num);
    int* my_nums = malloc(sizeof (int) * my_num);
    int win_cursor;
    int my_cursor;
    int curr;
    char c;

    slice_t slice = { 0 };

    fseek(fp, 0, SEEK_SET);
    while (true) {
        len = getline(&buf, &buf_size, fp);
        if (len == -1) {
            break;
        }

        state = 0;
        win_cursor = 0;
        my_cursor = 0;
        curr = -1;

        for (int i = 0; i < len; i++) {
            c = buf[i];
            switch (state) {
                case 0: {
                    if (c == ':') {
                        state = 1;
                    }
                    break;
                }
                case 1: {
                    if (c == '|') {
                        state = 2;
                    } else if (isdigit(c)) {
                        if (curr == -1) {
                            curr = c - 0x30;
                        } else {
                            curr *= 10;
                            curr += c - 0x30;
                        }
                    } else {
                        if (curr != -1) {
                            win_nums[win_cursor] = curr;
                            win_cursor += 1;
                            curr = -1;
                        }
                    }
                    break;
                }
                case 2: {
                    if (isdigit(c)) {
                        if (curr == -1) {
                            curr = c - 0x30;
                        } else {
                            curr *= 10;
                            curr += c - 0x30;
                        }
                    } else {
                        if (curr != -1) {
                            my_nums[my_cursor] = curr;
                            my_cursor += 1;
                            curr = -1;
                        }
                    }
                }
            }
        }

        int count = 0;
        for (int i = 0; i < win_num; i++) {
            for (int j = 0; j < my_num; j++) {
                if (win_nums[i] == my_nums[j]) {
                    count += 1;
                }
            }
        }
        append(&slice, count);
    }

    fclose(fp);
    free(buf);
    free(my_nums);
    free(win_nums);

    int sum = 0;
    // #pragma omp parallel for reduction(+: sum)
    for (int i = 0; i < slice.len; i++) {
        sum += check_card(&slice, i);
    }
    printf("part 2: %d\n", sum);

    free(slice.arr);

    return 0;
}
