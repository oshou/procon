#include <bits/stdc++.h>

int main() {
  int n, num;
  char symbol;
  scanf("%d", &n);
  bool s[13] = {false};
  bool h[13] = {false};
  bool c[13] = {false};
  bool d[13] = {false};
  for (int i = 0; i < n; i++) {
    scanf("%s %d", &symbol, &num);
    switch (symbol) {
      case 'S':
        s[num - 1] = true;
        break;
      case 'H':
        h[num - 1] = true;
        break;
      case 'C':
        c[num - 1] = true;
        break;
      case 'D':
        d[num - 1] = true;
        break;
    }
  }
  for (int i = 0; i < 13; i++) {
    if (!s[i]) {
      printf("S %d\n", i + 1);
    }
  }
  for (int i = 0; i < 13; i++) {
    if (!h[i]) {
      printf("H %d\n", i + 1);
    }
  }
  for (int i = 0; i < 13; i++) {
    if (!c[i]) {
      printf("C %d\n", i + 1);
    }
  }
  for (int i = 0; i < 13; i++) {
    if (!d[i]) {
      printf("D %d\n", i + 1);
    }
  }
}
