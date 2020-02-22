package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var x []int
	var sum int
	var min int = 1000000
	for i := 0; i < n; i++ {
		x = append(x, readInt())
	}
	sort.Ints(x)
	for p := x[0]; p <= x[len(x)-1]; p++ {
		sum = 0
		for _, v := range x {
			sum += (v - p) * (v - p)
		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min)
}
