#include <bits/stdc++.h>
using namespace std;
int main() {
  int a[3][3];
  for (int i = 0; i < 3; i++) {
    for (int j = 0; j < 3; j++) {
      cin >> a[i][j];
    }
  }
  int n;
  cin >> n;
  for (int k = 0; k < n; k++) {
    int hit;
    cin >> hit;
    for (int i = 0; i < 3; i++) {
      for (int j = 0; j < 3; j++) {
        if (a[i][j] == hit) {
          a[i][j] = 0;
        }
      }
    }
  }
  for (int i = 0; i < n; i++) {
    if ((a[i][0] == 0) && (a[i][1] == 0) && (a[i][2] == 0)) {
      cout << "Yes" << endl;
      return 0;
    }
  }
  for (int i = 0; i < n; i++) {
    if ((a[0][i] == 0) && (a[1][i] == 0) && (a[2][i] == 0)) {
      cout << "Yes" << endl;
      return 0;
    }
  }
  if (((a[0][0] == 0) && (a[1][1] == 0) && (a[2][2] == 0)) ||
      ((a[2][0] == 0) && (a[1][1] == 0) && (a[0][2] == 0))) {
    cout << "Yes" << endl;
    return 0;
  }
  cout << "No" << endl;
}
