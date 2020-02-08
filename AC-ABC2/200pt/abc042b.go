package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Strings(arr)
	fmt.Println(strings.Join(arr, ""))
}
