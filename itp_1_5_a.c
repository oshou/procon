/*
 * ITP_1_5_A
 *
 * input: The input consists of multiple datasets. Each dataset consists of two integers H and W separated by a single space.
 *        The input ends with two 0 (when both H and W are zero).
 * output: For each dataset, print the rectangle made of H × W '#'.
 *         Print a blank line after each dataset.
 */
#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main()
{
    int H,W;
    int h,w;
    while (1){
        scanf("%d %d", &H, &W);
        if (H == 0 && W == 0) break;
        for (h = 0; h < H; h++){
            for (w = 0; w < W; w++){
                putchar('#');
            }
            putchar('\n');
        }
    }
    return 0;
}
