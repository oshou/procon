package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	sorted := []int{a, b, c, d}
	sort.Ints(sorted)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	fmt.Println(10*sorted[0] + sorted[2] + 10*sorted[1] + sorted[3])
}
