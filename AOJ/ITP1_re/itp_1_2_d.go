package main

import "fmt"

func main() {
	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)
	if 0 <= x-r && x+r <= w && 0 <= y-r && y+r <= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
