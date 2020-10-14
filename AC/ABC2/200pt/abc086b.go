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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b := readString(), readString()
	num, _ := strconv.ParseFloat(a+b, math.MaxInt64)
	sqrt := math.Sqrt(num)
	if math.Floor(sqrt) == sqrt {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
