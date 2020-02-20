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

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a, b := readInt(), readInt()
	switch {
	case a < b:
		fmt.Println("a < b")
	case a > b:
		fmt.Println("a > b")
	default:
		fmt.Println("a == b")
	}
}
