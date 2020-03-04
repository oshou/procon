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

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	n := readInt()
	s := make([]bool, 14)
	h := make([]bool, 14)
	c := make([]bool, 14)
	d := make([]bool, 14)
	for i := 0; i < n; i++ {
		symbol, num := readString(), readInt()
		switch symbol {
		case "S":
			s[num] = true
		case "H":
			h[num] = true
		case "C":
			c[num] = true
		case "D":
			d[num] = true
		}
	}
	for i := 1; i <= 13; i++ {
		if !s[i] {
			fmt.Printf("S %d\n", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !h[i] {
			fmt.Printf("H %d\n", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !c[i] {
			fmt.Printf("C %d\n", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !d[i] {
			fmt.Printf("D %d\n", i)
		}
	}
}
