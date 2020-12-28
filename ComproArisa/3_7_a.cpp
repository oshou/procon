#include <bits/stdc++.h>
using namespace std;

int main() {
  int n;
  cin >> n;
  string s;
  cin >> s;
  vector<int> vE(n + 1, 0);
  vector<int> vW(n + 1, 0);
  for (int i = 0; i < n; i++) {
    if (i >= 1) vE[i + 1] += vE[i];
    if (s[i] == "E") vE[i + 1] += 1;
    if (i >= 1) vW[i + 1] += vW[i];
    if (s[i] == "W") vW[i] += 1;
  }
}
