package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var memo []int

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func factorial(n int) int {
	if memo[n-1] != 0 {
		return memo[n-1]
	}
	if n == 1 {
		memo[0] = 1
		return memo[0]
	}
	memo[n-1] = n * factorial(n-1) % 1000000007
	return memo[n-1]
}

func combination(n, k int) int {
	return factorial(n)
	//return (factorial(n) / (factorial(k) * factorial(k))) % 1000000007
}

func main() {
	n, a, b := readInt(), readInt(), readInt()
	sum := 0
	memo = make([]int, n)
	for k := 2; k <= n; k++ {
		if k == a || k == b {
			continue
		}
		sum += combination(n, k)
	}
	fmt.Println(sum)
}
