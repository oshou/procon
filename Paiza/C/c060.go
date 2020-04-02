package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var n, k, p int
	var arr []string
	fmt.Scan(&n, &k, &p)
	sc.Scan()
	arr = make([]string, n)
	arr = strings.Split(sc.Text(), " ")
	sort.Strings(arr)
	for i := k * (p - 1); i < min(len(arr), k*(p-1)+k); i++ {
		fmt.Println(arr[i])
	}
}
