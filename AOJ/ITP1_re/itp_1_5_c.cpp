#include <bits/stdc++.h>

int main() {
  int h, w, remi, remj;
  while (true) {
    scanf("%d %d", &h, &w);
    if (h == 0 && w == 0) {
      break;
    }
    for (int i = 1; i <= h; i++) {
      remi = i % 2;
      for (int j = 1; j <= w; j++) {
        remj = j % 2;
        if ((remi == 0 && remj == 0) || (remi == 1 && remj == 1)) {
          printf("#");
        } else {
          printf(".");
        }
      }
      printf("\n");
    }
    printf("\n");
  }
}
