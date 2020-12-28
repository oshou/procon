package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i > 9 {
			fmt.Println("No")
			return
		}
		if i*i > n {
			fmt.Println("No")
			return
		}
		var res = n / i
		if n%i == 0 && 1 <= res && res <= 9 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
