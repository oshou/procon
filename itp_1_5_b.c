/*
 * ITP_1_5_B
 *
 * input:
 * The input consists of multiple datasets. Each dataset consists of two integers H and W separated by a single space.
 * The input ends with two 0 (when both H and W are zero).
 *
 * output:
 * For each dataset, print the frame made of '#' and '.'.
 * Print a blank line after each dataset.
 */
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

int main()
{
    int H,W;
    int h,w;
    while (1){
        scanf("%d %d", &H, &W);
        if (H == 0 && W == 0) break;
        for (h = 1; h <= H; h++){
            for (w = 1; w <= W; w++){
                if ( h != 1 && h != H && w != 1 && w != W){
                    putchar('.');
                } else {
                    putchar('#');
                }
            }
            putchar('\n');
        }
    }
    return 0;
}
