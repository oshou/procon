#include <stdio.h>

int main()
{
    int a,b = 1,c;
    while(1){
        scanf("%d",&a);
        if(a == 0) break;
        printf("Case %d: %d\n", b, a);
        b++;
    }
    return 0;
}
