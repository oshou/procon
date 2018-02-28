/*
 * ITP_1_6_A
 *
 * input:
 * The input is given in the following format:
 * n
 * a1 a2 . . . an
 *
 * n is the size of the sequence and ai is the ith element of the sequence.
 *
 * output:
 *
 * Print the reversed sequence in a line.
 * Print a single space character between adjacent elements
 * (Note that your program should not put a space character after the last element).
 */
#include <stdio.h>

int main()
{
    int i, n;
    short a[101];

    scanf("%d", &n);
    for (i = 0; i < n; i++) scanf("%d", &a[i]);
    if (n >= 1) printf("%d", a[n-1]);
    for (i = n - 2; i >= 0; i--) printf(" %d", a[i]);
    putchar('\n');
    return 0;
}
