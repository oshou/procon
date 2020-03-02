#include <bits/stdc++.h>
using namespace std;
int main() {
  int a, b;
  double a_f, b_f;
  scanf("%d %d", &a, &b);
  a_f = static_cast<double>(a);
  b_f = static_cast<double>(b);
  printf("%d %d %.5f\n", a / b, a % b, a_f / b_f);
}
