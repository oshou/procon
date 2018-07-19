/*
 * ITP_1_4_D
 *
 * input:
 * In the first line, an integer n is given. In the next line, n integers ai are given in a line.
 *
 * output:
 * Print the minimum value, maximum value and sum in a line. Put a single space between the values.
 */
#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main()
{
    int n;
    int d, min, max;
    long long sum;

    scanf("%d%d", &n, &d);
    sum = min = max = d;
    while (--n > 0) {
        scanf("%d", &d);
        if (d < min) min = d;
        if (d > max) max = d;
        sum += d;
    }
    printf("%d %d %lld\n", min, max, sum);
    return 0;
}
