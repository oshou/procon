#include <bits/stdc++.h>
using namespace std;

int main() {
  string s;
  cin >> s;
  int n = 0;
  for (int i = 0; i < s.size(); i++) {
    n += s[i] - '0';
  }
  if (stoi(s) % n == 0) {
    cout << "No" << endl;
  } else {
    cout << "Yes" << endl;
  }
}
