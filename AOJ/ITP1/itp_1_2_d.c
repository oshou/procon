#include <stdio.h>
#include <stdlib.h>

int main()
{
    int W, H, x, y, r;
    scanf("%d %d %d %d %d", &W, &H, &x, &y, &r);

    if (r <= x && r <= y && (H - r) >= y && (W - r) >= x){
        printf("Yes\n");
    } else {
        printf("No\n");
    }
    return 0;
}
