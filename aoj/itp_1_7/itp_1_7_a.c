#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main()
{
    int m, f, r, sum;

    while (1){
        scanf("%d%d%d", &m, &f, &r);
        sum = m + f;

        if (m == -1 && f == -1 && r == -1) {
            break;
        }

        if (m < 0 || f < 0) {
            puts("F");
        } else if (sum >= 80) {
            puts("A");
        } else if (sum >= 65) {
            puts("B");
        } else if (sum >= 50) {
            puts("C");
        } else if (sum >= 30){
            if (r >= 50){
                puts("C");
            } else {
                puts("D");
            }
        } else {
            puts("F");
        }
    }
}
