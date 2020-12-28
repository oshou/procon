#include <bits/stdc++.h>
using namespace std;

int main() {
  int ans = 0;
  int n;
  cin >> n;
  vector<int> a(n);
  bool div_by_2 = true;
  while (div_by_2) {
    for (int i = 0; i < n; i++) {
      cin >> a[i];
      if (a[i] % 2 == 1) {
        div_by_2 = false;
        break;
      } else {
        a[i] /= 2;
      }
    }
    if (div_by_2) {
      ans++;
    }
  }
  cout << ans << endl;
}
