#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  scanf("%d", &n);
  int tmp;
  scanf("%d", &tmp);
  cout << tmp;
  int sum = 0;
  int min = n == 0 ? 0 : 1000000;
  int max = n == 0 ? 0 : -1000000;
  for (int i = 0; i < n; i++) {
    scanf("%d", &tmp);
    sum += tmp;
    if (min > tmp) min = tmp;
    if (max < tmp) max = tmp;
  }
  printf("%d %d %d\n", min, max, sum);
}
