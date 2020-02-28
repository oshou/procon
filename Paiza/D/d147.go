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
	s := readInt()
	a, b, c, d := readInt(), readInt(), readInt(), readInt()
	min := min(min(min(a, b), c), d)
	fmt.Println(s * min)
}
