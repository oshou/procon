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

func main() {
	var cnt int
	n, m, k := readInt(), readInt(), readInt()
	// 友達関係
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = readInt(), readInt()
	}
	// ブロック関係
	c := make([]int, k)
	d := make([]int, k)
	for i := 0; i < k; i++ {
		c[i], d[i] = readInt(), readInt()
	}

	for i := 0; i < n; i++ {
		cnt = 0
		for j := 0; j < n; j++ {
			// is-friend
			for l := 0; l < m; l++ {
				if a[l] == i && b[l] == j {
					continue
				}
			}
			// is-blocked
			for l := 0; l < k; l++ {
				if c[l] == i && d[l] == j {
					continue
				}
			}
			// is-indirect-friend
		}
		fmt.Printf("%v ", cnt)
	}
	fmt.Println()
}
