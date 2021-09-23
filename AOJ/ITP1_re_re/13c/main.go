package main

import "fmt"

func swap(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	var (
		a, b       int
		arra, arrb []int
	)
	for {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		arra = append(arra, a)
		arrb = append(arrb, b)
	}
	for i := 0; i < len(arra); i++ {
		if arra[i] > arrb[i] {
			fmt.Println(swap(arra[i], arrb[i]))
		} else {
			fmt.Println(arra[i], arrb[i])
		}
	}
}
