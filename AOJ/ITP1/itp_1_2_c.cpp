#include <bits/stdc++.h>
using namespace std;
int main() {
  int arr[3];
  for (int i = 0; i < 3; i++) {
    cin >> arr[i];
  }
  for (int i = 2; i > 0; i--) {
    for (int j = 0; j < i; j++) {
      if (arr[j] > arr[j + 1]) {
        swap(arr[j], arr[j + 1]);
      }
    }
  }
  printf("%d %d %d\n", arr[0], arr[1], arr[2]);
}
