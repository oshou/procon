package main

import "fmt"

func main() {
	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)
	if (x-r >= 0) && (y-r >= 0) && (x+r <= w) && (y+r <= h) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
