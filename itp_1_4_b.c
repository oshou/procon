/*
 * input: a op b
 * output: for each dataset, print the value line
 */
#include <stdio.h>

int main()
{
    int i,a,b;
    char op;
    for (i=1;i>0;i++){
        scanf("%d %c %d",&a,&op,&b);
        if ( op == '?')
            break;
        else if ( op == '+')
            printf("%d\n",a+b);
        else if ( op == '-')
            printf("%d\n",a-b);
        else if ( op == '*')
            printf("%d\n",a*b);
        else if ( op == '/')
            printf("%d\n",a/b);
    }
    return 0;
}
