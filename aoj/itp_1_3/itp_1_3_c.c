#include <stdio.h>

int main()
{
    int i,n,m;
    for (i=1;;i++){
        scanf("%d %d",&n,&m);
        if (n==0 && m==0) break;
        if (n<m){
            printf("%d %d",&n,&m);
        } else {
            printf("%d %d",&m,&n);
        }
    }
}
