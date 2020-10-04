package main

import "fmt"

func main() {
	// s*t >= d ならok
	// O(1)
	var d, t, s int
	fmt.Scan(&d, &t, &s)
	if s*t >= d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// package a
//
// func IsInTime(d, t, s int) bool {
// 	if s*t >= d {
// 		return true
// 	} else {
// 		return false
// 	}
// }
