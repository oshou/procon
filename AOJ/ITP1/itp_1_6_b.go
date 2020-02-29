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
	s := make([]bool, 13)
	h := make([]bool, 13)
	c := make([]bool, 13)
	d := make([]bool, 13)
	for i := 0; i < n; i++ {
		symbol, num := readString(), readInt()
		switch symbol {
		case "S":
			s[num-1] = true
		case "H":
			h[num-1] = true
		case "C":
			c[num-1] = true
		case "D":
			d[num-1] = true
		}
	}
	for i := 0; i < 13; i++ {
		if s[i] == false {
			fmt.Println("S", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if h[i] == false {
			fmt.Println("H", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if c[i] == false {
			fmt.Println("C", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if d[i] == false {
			fmt.Println("D", i+1)
		}
	}
}
