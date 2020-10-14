package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	m := make(map[int]int)
	keys := []int{1}
	for i := 2; i <= n; i++ {
		fmt.Scan(&a)
		m[a]++
		keys = append(keys, i)
	}
	for _, key := range keys {
		fmt.Println(m[key])
	}
}
