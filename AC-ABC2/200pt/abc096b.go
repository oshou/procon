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

func readFloat() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), math.MaxInt64)
	return r
}

func main() {
	a, b, c := readFloat(), readFloat(), readFloat()
	k := readFloat()
	max := math.Max(math.Max(a, b), c)
	sum := max*math.Pow(2, k) + (a + b + c) - max
	fmt.Println(sum)
}
