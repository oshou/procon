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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	h, _, x := readInt(), readInt(), readInt()
	var s string
	for i := 0; i < h; i++ {
		a := readString()
		s += a
	}
	for i := 0; i < len(s); i += x {
		fmt.Println(string(s[i:min(i+x, len(s))]))
	}
}
