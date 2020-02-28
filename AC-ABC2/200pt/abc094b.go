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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	_, m, x := readInt(), readInt(), readInt()
	var a []int
	for i := 0; i < m; i++ {
		a = append(a, readInt())
	}
	for i := 0; i < m; i++ {
		if a[i] > x {
			fmt.Println(min(len(a[:i]), len(a[i:])))
			break
		}
	}
}
