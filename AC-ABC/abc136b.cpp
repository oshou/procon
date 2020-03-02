#include <bits/stdc++.h>
using namespace std;

int main() {
  int n, cnt = 0;
  cin >> n;
  for (int i = 1; i <= n; i++) {
    if (to_string(i).length() % 2 != 0) {
      cnt++;
    }
  }
  cout << cnt << endl;
}
