#include <bits/stdc++.h>
using namespace std;
int main() {
  int n, tmp;
  cin >> n;
  for (int i = 1; i <= n; i++) {
    if (i % 3 == 0) {
      cout << " " << i;
      continue;
    }
    tmp = i;
    while (tmp > 0) {
      if (tmp % 10 == 3) {
        cout << " " << i;
        break;
      }
      tmp /= 10;
    }
  }
  cout << endl;
}
