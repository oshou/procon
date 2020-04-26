package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		txt, _ := strconv.Atoi(sc.Text())
		arr[i] = txt
	}
	sort.Sort(sort.IntSlice(arr))

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
