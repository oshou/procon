#include <bits/stdc++.h>
using namespace std;
int main() {
  int n, tmp;
  cin >> n;
  char ch;
  bool s[14], h[14], c[14], d[14];
  for (int i = 0; i < n; i++) {
    cin >> ch >> tmp;
    switch (ch) {
      case 'S':
        s[tmp] = true;
        break;
      case 'H':
        h[tmp] = true;
        break;
      case 'C':
        c[tmp] = true;
        break;
      case 'D':
        d[tmp] = true;
        break;
    }
  }
  for (int i = 1; i <= 13; i++) {
    if (!s[i]) printf("S %d\n", i);
  }
  for (int i = 1; i <= 13; i++) {
    if (!h[i]) printf("H %d\n", i);
  }
  for (int i = 1; i <= 13; i++) {
    if (!c[i]) printf("C %d\n", i);
  }
  for (int i = 1; i <= 13; i++) {
    if (!d[i]) printf("D %d\n", i);
  }
}
