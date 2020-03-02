#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  cin >> n;
  int h, m, s;
  h = n / 3600;
  m = (n % 3600) / 60;
  s = (n % 3600) % 60;
  printf("%d:%d:%d\n", h, m, s);
}
