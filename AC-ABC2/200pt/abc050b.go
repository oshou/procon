package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// Input
	var n, m int
	n = readInt()
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = readInt()
	}
	m = readInt()
	p := make([]int, m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		p[i], x[i] = readInt(), readInt()
	}

	// Calculate
	sum := sum(t)
	for i := 0; i < m; i++ {
		fmt.Println(sum - t[p[i]-1] + x[i])
	}
}
