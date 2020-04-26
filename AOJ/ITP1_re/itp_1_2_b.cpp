#include <bits/stdc++.h>
using namespace std;

int main() {
  int a, b, c;
  scanf("%d %d %d", &a, &b, &c);
  if (a < b && b < c) {
    cout << "Yes" << endl;
  } else {
    cout << "No" << endl;
  }
}
