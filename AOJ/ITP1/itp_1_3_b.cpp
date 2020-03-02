#include <bits/stdc++.h>
using namespace std;
int main() {
  int i, v;
  i = 0;
  while (true) {
    cin >> v;
    if (v == 0) {
      return 0;
    }
    i++;
    printf("Case %d: %d\n", i, v);
  }
}
