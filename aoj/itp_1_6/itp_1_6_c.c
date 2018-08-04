/**
 * Input:
 * In the first line, the number of notices n is given.
 * In the following n lines, a set of four integers b, f, r and v which represents ith notice is given in a line.
 *
 * Output:
 * For each building, print the information of 1st, 2nd and 3rd floor in this order.
 * For each floor information, print the number of tenants of 1st, 2nd, .. and 10th room in this order.
 * Print a single space character before the number of tenants.
 * Print "####################" (20 '#') between buildings.
 */

#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main(){
    /**
     * n: entry
     *
     * b: buildings (max:4)
     * f: floors (max:3)
     * r: rooms (max:10)
     * v: person
     */
    int i, n;
    int b, f, r, v;
    int house[5][4][11];

    // Initialize
    for (b = 1; b <= 4; b++){
        for (f = 1; f <= 3; f++){
            for (r = 1; r <= 10; r++) house[b][f][r] = 0;
        }
    }

    // Input
    scanf("%d", &n);
    for (i = 0; i < n; i++){
        scanf("%d%d%d%d", &b, &f, &r, &v);
        // house[b][f][r] = house[b][f][r] + v;
        house[b][f][r] += v;
    }

    // Output
    for (b = 1; b <= 4; b++){
        for (f = 1; f <= 3; f++){
            for (r = 1; r <= 10; r++){
                printf(" %d", house[b][f][r]);
            }
            putchar('\n');
        }
        if (b < 4) puts("####################");
    }
}
