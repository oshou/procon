#include <bits/stdc++.h>
using namespace std;
int main() {
  int h, w;
  bool iEven, jEven;
  while (true) {
    scanf("%d %d", &h, &w);
    if ((h == 0) && (w == 0)) break;
    for (int i = 0; i < h; i++) {
      for (int j = 0; j < w; j++) {
        iEven = i % 2 == 0;
        jEven = j % 2 == 0;
        if (iEven && jEven) {
          printf("#");
        } else if (iEven && !jEven) {
          printf(".");
        } else if (!iEven && jEven) {
          printf(".");
        } else {
          printf("#");
        }
      }
      printf("\n");
    }
    printf("\n");
  }
}
