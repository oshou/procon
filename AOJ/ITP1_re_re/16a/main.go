package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	arr []int
)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	for i := n - 1; i >= 0; i-- {
		if i != n-1 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()
}
