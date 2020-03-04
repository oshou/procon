#include <bits/stdc++.h>
using namespace std;
int main() {
  string s;
  getline(cin, s);
  for (int i = 0; i < s.length(); i++) {
    if (isupper(s[i])) {
      cout << tolower(s[i]);
    } else {
      cout << toupper(s[i]);
    }
  }
}
