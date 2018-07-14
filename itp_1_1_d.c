#include <stdio.h>
#include <stdlib.h>

int main()
{
    int sec;
    scanf("%d",&sec);
    printf("%d:%d:%d\n",sec / (60 * 60),sec % (60 * 60) / 60,sec % (60 * 60) % 60);
    return 0;
}
