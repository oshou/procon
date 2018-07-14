/*
 * ITP_1_6_C
 *
 * input:
 * v: person
 * r: room
 * f: fth room number
 * b: building
 */

#include <stdio.h>

int main()
{
    int n;
    int room[10][3][4];
    int i, j, k;
    int a, b, c, d;
    for (i=0; i<10; i++){
        for (j=0; j<3; j++){
            for (k=0; k<4; k++){
                room[i][j][k]=0;
            }
        }
    }
    scanf("%d", &n);
    for (i=0; i<n; i++){
        scanf("%d %d %d %d", &a, &b, &c, &d);
        room[c-1][b-1][a-1]+=d;
    }
    for (i=0; i<4; i++){
        for (j=0; j<3; j++){
            for (k=0; k<10; k++){
                printf(" ");
                printf("%d", room[k][j][i]);
            }
            puts("");
        }
        if (i==3) break;
        for (j=0; j<20; j++) printf("#");
        puts("");
    }
    return 0;
}

