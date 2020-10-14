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

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func sum(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	n, k := readInt(), readInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = readInt()
	}
	var s, max int = -1, -1
	var exp float64
	for i := 0; i <= n-k; i++ {
		s = sum(p[i : i+k]...)
		if s > max {
			max = s
			exp = float64(s+k) / 2
		}
	}
	fmt.Println(exp)
}
