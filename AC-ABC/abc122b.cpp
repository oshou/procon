#include <bits/stdc++.h>
using namespace std;

int main() {
  string s;
  cin >> s;
  int max, tmp = 0;
  for (int i = 0; i < s.length(); i++) {
    switch (s[i]) {
      case 'A':
      case 'C':
      case 'G':
      case 'T':
        tmp++;
        break;
      default:
        if (max < tmp) {
          max = tmp;
          tmp = 0;
        }
    }
    cout << "tmp:" << tmp << endl;
  }
  if (max < tmp) {
    max = tmp;
    tmp = 0;
  }
  cout << max << endl;
}
