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
	d, x := readInt(), readInt()
	var a []int
	var sum int
	for i := 0; i < n; i++ {
		a = append(a, readInt())
	}

	for i := 0; i < n; i++ {
		sum += 1 + (d-1)/a[i]
	}
	fmt.Println(sum + x)
}
