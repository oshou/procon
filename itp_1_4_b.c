/*
 * input: r(real)
 * output: Print the area and circumference of the circle in a line. Put a single space between them. The output should not contain an absolute error greater than 10-5.
 */
#include <stdio.h>
#include <stdlib.h>

#define PI 3.1415926535897932

int main()
{
    double r;
    scanf("%lf",&r);
    printf("%lf %lf\n", PI*r*r, 2*PI*r);
    return 0;
}
