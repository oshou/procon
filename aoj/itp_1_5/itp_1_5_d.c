/*
 * ITP_1_5_D
 *
 * input:
 * An integer n is given in a line.
 *
 * output:
 * Print the output result of the above program for given integer n.
 */
void call(int n)
{
    int i = 1;
    CHECK_NUM;
    int x = i;
    if (x % 3 == 0){
        cout << " " << i;
        goto END_CHECK_NUM;
    }
    INCLUDE 3:
        if (x % 10 == 3){
            cout << " " << i;
            goto END_CHECK_NUM;
        }
    x /= 10;
    if ( x ) goto INCLUDE 3;
    END_CHECK_NUM;
    if (++i <= n) goto CHECK_NUM;

    cout << endl;
}


void call(int n)
{
    int i = 1;
    int x;

    do {
        x = i;
        if (x % 3 == 0){
            printf(" %d", i);
        }
        else {
            do {
                if (x % 10 == 3){
                    printf(" %d", i);
                    x = 0;
                }
                else x /= 10;
            } while (x);
        }
    } while (++i <= n);
    putchar('\n');
}
