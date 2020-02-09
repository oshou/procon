package main

import (
	"bufio"
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
	a, b, c := readInt(), readInt(), readInt()
	(a*n)%b == c

}

a-b

2*n%5 == 1

99*n%100 == 97
