#include <bits/stdc++.h>

int main() {
  int x, y;
  while (true) {
    scanf("%d %d", &x, &y);
    if (x == 0 && y == 0) {
      break;
    }
    if (x > y) {
      printf("%d %d\n", y, x);
    } else {
      printf("%d %d\n", x, y);
    }
  }
}
