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
	n := readInt()
	l := make([]int, n+1)
	l[0], l[1] = 2, 1
	for i := 2; i <= n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	fmt.Println(l[n])
}
