#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  int tmp;
  int sum = 0;
  int min = 1000000;
  int max = -1000000;
  scanf("%d", &n);
  for (int i = 0; i < n; i++) {
    scanf("%d", &tmp);
    sum += tmp;
    if (min > tmp) min = tmp;
    if (max < tmp) max = tmp;
  }
  printf("%d %d %d\n", min, max, sum);
}
