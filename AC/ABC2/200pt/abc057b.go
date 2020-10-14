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

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func main() {
	n, m := readInt(), readInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = readInt(), readInt()
	}
	c := make([]int, m)
	d := make([]int, m)
	for i := 0; i < m; i++ {
		c[i], d[i] = readInt(), readInt()
	}

	for i := 0; i < n; i++ {
		min := -1
		pos := -1
		for j := 0; j < m; j++ {
			d := abs(a[i]-c[j]) + abs(b[i]-d[j])
			if min == -1 || d < min {
				min = d
				pos = j + 1
			}
		}
		fmt.Println(pos)
	}
}
