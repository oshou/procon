#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  int b, f, r, v;
  cin >> n;
  int room[4][3][10];
  for (int i = 0; i < 4; i++) {
    for (int j = 0; j < 3; j++) {
      for (int k = 0; k < 10; k++) {
        room[i][j][k] = 0;
      }
    }
  }
  for (int i = 0; i < n; i++) {
    cin >> b >> f >> r >> v;
    room[b - 1][f - 1][r - 1] = max(room[b - 1][f - 1][r - 1] + v, 0);
  }
  for (int i = 0; i < 4; i++) {
    if (i != 0) {
      cout << "####################" << endl;
    }
    for (int j = 0; j < 3; j++) {
      for (int k = 0; k < 10; k++) {
        cout << " " << room[i][j][k];
      }
      cout << endl;
    }
  }
}
