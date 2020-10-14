#include <stdio.h>
using namespace std;

int main() {
  int n;
  int j;
  scanf("%d", &n);
  for (int i = 1; i <= 9; i++) {
    for (int j = 1; j <= 9; j++) {
      if (i * j == n) {
        printf("Yes\n");
        return 0;
      }
    }
  }
  printf("No\n");
  return 0;
}
