#include <bits/stdc++.h>

int main() {
  int h, w;
  while (true) {
    scanf("%d %d", &h, &w);
    if (h == 0 && w == 0) {
      break;
    }
    for (int i = 1; i <= h; i++) {
      for (int j = 1; j <= w; j++) {
        if (i == 1 || i == h) {
          printf("#");
          continue;
        }
        if (j == 1 || j == w) {
          printf("#");
          continue;
        }
        printf(".");
      }
      printf("\n");
    }
    printf("\n");
  }
}
