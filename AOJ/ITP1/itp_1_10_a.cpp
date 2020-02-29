#include <stdio.h>
#include <cmath>

using namespace std;

int main() {
  double x1, y1, x2, y2;
  scanf("%lf %lf %lf %lf", &x1, &y1, &x2, &y2);
  printf("%.8f\n", sqrt(pow((x2 - x1), 2) + pow((y2 - y1), 2)));
}
