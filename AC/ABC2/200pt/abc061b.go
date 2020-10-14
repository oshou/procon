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
	n, m := readInt(), readInt()
	a := make([]int, m)
	b := make([]int, m)
	hash := make([]int, n)
	for i := 0; i < m; i++ {
		a[i], b[i] = readInt(), readInt()
		hash[a[i]-1]++
		hash[b[i]-1]++
	}
	for _, v := range hash {
		fmt.Println(v)
	}
}
