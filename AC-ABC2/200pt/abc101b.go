package main

import "fmt"

func digits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var n int
	fmt.Scan(&n)
	if n%digits(n) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
