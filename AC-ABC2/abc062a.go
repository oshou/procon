package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	m := map[int]int{1: 1, 2: 3, 3: 1, 4: 2, 5: 1, 6: 2, 7: 1, 8: 1, 9: 2, 10: 1, 11: 2, 12: 1}
	if m[x] == m[y] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
