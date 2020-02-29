#include <stdio.h>

using namespace std;

int main() {
  int m, f, r;
  char grade;
  int score;
  while (true) {
    scanf("%d %d %d", &m, &f, &r);
    score = m + f;
    if (m == -1 && f == -1 && r == -1) break;
    if (m == -1 || f == -1) {
      grade = 'F';
    } else if (80 <= score) {
      grade = 'A';
    } else if (65 <= score) {
      grade = 'B';
    } else if (50 <= score || 50 <= r) {
      grade = 'C';
    } else if (30 <= score) {
      grade = 'D';
    } else {
      grade = 'F';
    }
    printf("%c\n", grade);
  }
  return 0;
}
