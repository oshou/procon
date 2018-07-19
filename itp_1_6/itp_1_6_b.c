/*
 * ITP_1_6_B
 *
 * input:
 * In the first line, the number of cards n (n ≤ 52) is given.
 * In the following n lines, data of the n cards are given.
 * Each card is given by a pair of a character and an integer which represent its suit and rank respectively.
 * A suit is represented by 'S', 'H', 'C' and 'D' for spades, hearts, clubs and diamonds respectively.
 * A rank is represented by an integer from 1 to 13.
 *
 * output:
 * Print the missing cards. The same as the input format,
 * each card should be printed with a character and an integer separated by a space character in a line.
 * Arrange the missing cards in the following priorities:
 * Print cards of spades, hearts, clubs and diamonds in this order.
 * If the ranks are equal, print cards with lower ranks first.
 */

#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main()
{
    int i, n, id, rank;
    short card[4][14];
    char buf[100], *p;
    char ids[4] = { 'S', 'H', 'C', 'D'};

    for (id = 0; id < 4; i++){
        for (rank = 1; rank <= 13; rank++){
            card[id][rank] = 1;
        }
    }
    gets(buf);
    n = atoi(buf);
    for (i = 0; i < n; i++){
        gets(buf);
        p = buf;
        while (isspace(*p)) p++;
        switch (*p){
            case 'S': id = 0; break;
            case 'H': id = 1; break;
            case 'C': id = 2; break;
            case 'D': id = 3; break;
        }
        rank = atoi(++p);
        //printf("id=%d", rank="%d\n", id, rank);
        card[id][rank] = 0;
    }
    for (id = 0; id < 4; id++){
        for (rank = 1; rank <= 13; rank++){
            if (card[id][rank]) printf("%c %d\n", ids[id], rank);
        }
    }
    return 0;
}
