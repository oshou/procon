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

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func main() {
	a, b := readFloat64(), readFloat64()
	switch {
	case a == b:
		fmt.Println("EQUAL")
	case a > b:
		fmt.Println("GREATER")
	case a < b:
		fmt.Println("LESS")
	}
}
