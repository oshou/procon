package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			return
		}
		fmt.Println(strings.Repeat("#", w))
		for i := 0; i < h-2; i++ {
			fmt.Printf("#%v#\n", strings.Repeat(".", w-2))
		}
		fmt.Println(strings.Repeat("#", w))
		fmt.Println()
	}
}

#include <bits/stdc++.h>
using namespace std;
int main() {
  int n;
  int a[n];
  cin >> n;
  for (int i=0;i<n;i++){
    cin >> a[i];
  }
  for (int i=n-1;i>=0;i--){
    if (i != n-1){
      cout << " ";
    }
    cout << a[i];
  }
  cout << endl;
}
