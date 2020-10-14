#include <bits/stdc++.h>
using namespace std;
int main() {
  int n, cnt, num = 0;
  cin >> n;
  for (int i = 1; i <= n; i += 2) {
    cnt = 0;
    for (int j = 1; j <= n; j++) {
      if (i % j == 0) {
        cnt++;
      }
    }
    if (cnt == 8) {
      num++;
    }
  }
  cout << num << endl;
}
