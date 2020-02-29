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
}

func readFloat64() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), math.MaxInt64)
	return r
}

func main() {
	r := readFloat64()
	fmt.Printf("%05f %05f\n", r*r*math.Pi, 2*r*math.Pi)
}
