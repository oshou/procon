package main

import "fmt"

func less(e []int, i, j int) int {
	if e[i-1] > e[j-1] {
		return j
	}
	return i
}

func main() {
	var p1, p2, p3, p4 int
	fmt.Scan(&p1, &p2, &p3, &p4)

	e := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&e[i])
	}

	win1 := less(e, p1, p2)
	win2 := less(e, p3, p4)
	ans := make([]int, 2)
	if win1 < win2 {
		ans[0] = win1
		ans[1] = win2
	} else {
		ans[0] = win2
		ans[1] = win1
	}

	var f1, f2 int
	fmt.Scan(&f1, &f2)
	if f1 < f2 {
		fmt.Println(ans[0])
		fmt.Println(ans[1])
	} else {
		fmt.Println(ans[1])
		fmt.Println(ans[0])
	}
}
