#include <bits/stdc++.h>

int main() {
  int n, tmp;
  bool exists;
  scanf("%d", &n);
  for (int i = 1; i <= n; i++) {
    exists = false;
    tmp = i;
    if (tmp % 3 == 0) {
      printf(" %d", i);
      continue;
    }
    while (tmp != 0) {
      if (tmp % 10 == 3) {
        exists = true;
      }
      tmp /= 10;
    }
    if (exists) {
      printf(" %d", i);
      continue;
    }
  }
  printf("\n");
}
