package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var a, b, x, x0, n int
	var tak, tak_odd, tak_even int
	var aoki, aoki_odd, aoki_even int
	fmt.Scan(&a, &b, &x0, &n)
	x = x0
	for i := 0; i < n*2; i++ {
		x = (a*x + b) % 6
		if i%2 == 0 {
			tak += x%6 + 1
			fmt.Println("tak:", tak)
			if (tak)%2 == 0 {
				tak_even++
			} else {
				tak_odd++
			}
		} else {
			aoki += x + 1
			fmt.Println("aoki:", aoki)
			if (aoki)%2 == 0 {
				aoki_even++
			} else {
				aoki_odd++
			}
		}
	}
	fmt.Println("tak_even:", tak_even)
	fmt.Println("tak_odd:", tak_odd)
	fmt.Println("aoki_even:", aoki_even)
	fmt.Println("aoki_odd:", aoki_odd)
	fmt.Println(min(tak_even, tak_odd), min(aoki_even, aoki_odd))
}
