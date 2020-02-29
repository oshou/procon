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
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func digits(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	return sum
}

func main() {
	var s string
	for {
		s = readString()
		if s == "0" {
			break
		}
		fmt.Println(digits(s))
	}
}
