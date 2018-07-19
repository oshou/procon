#include <stdio.h>
#include <stdlib.h>

int main()
{
    int a, b, c;
    int k;
    int n;
    scanf("%d%d%d", &a, &b, &c);
    n = 0;
    for (k = a; k <= b; k++){
        if (c % k == 0) n++;
    }
    printf("%d\n", n);
    return 0;
}
