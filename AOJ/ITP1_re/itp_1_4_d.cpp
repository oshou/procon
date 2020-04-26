#include <bits/stdc++.h>

int main() {
  long n, a, min, max, sum;
  scanf("%ld", &n);
  for (int i = 1; i <= n; i++) {
    scanf("%ld", &a);
    if (i == 1) {
      min = a;
      max = a;
      sum = a;
      continue;
    }
    if (a < min) {
      min = a;
    }
    if (max < a) {
      max = a;
    }
    sum += a;
  }
  printf("%ld %ld %ld\n", min, max, sum);
}
