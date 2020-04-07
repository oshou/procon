package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	n := readInt()
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, 60)
	for i := 0; i < n; i++ {
		a[i], b[i] = readInt(), readInt()
		c[b[i]+a[i]-1]++
	}
	max := 0
	maxDay := 0
	for i, v := range c {
		if v > max {
			max = v
			maxDay = i + 1
		}
	}
	fmt.Println(maxDay)
}
