package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m, f, s int
	var arr [][]string
	fmt.Scan(&n)
	arr = make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
		arr[i][i] = "-"
	}

	m = n * (n - 1) / 2
	for i := 0; i < m; i++ {
		fmt.Scan(&f, &s)
		arr[f-1][s-1] = "W"
		arr[s-1][f-1] = "L"
	}

	for i := 0; i < n; i++ {
		fmt.Println(strings.Join(arr[i], " "))
	}
}
