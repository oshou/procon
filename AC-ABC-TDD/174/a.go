package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x >= 30 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

//package a
//
//func NeedsAirConditioner(x int) bool {
//	// x >= 30だったらtrue, 違うならfalse
//	// O(1) 定数時間で終わる
//	if x >= 30 {
//		return true
//	}else{
//
//	return false
//}
//}
