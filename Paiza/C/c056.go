package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n, m := readInt(), readInt()
	var a, b int
	for i := 1; i <= n; i++ {
		a, b = readInt(), readInt()
		if max(0, a-5*b) >= m {
			fmt.Println(i)
		}
	}
}
