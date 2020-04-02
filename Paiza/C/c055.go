package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	n := readInt()
	g := readString()
	var s string
	var isNone bool = true
	for i := 0; i < n; i++ {
		s = readString()
		if strings.Contains(s, g) {
			isNone = false
			fmt.Println(s)
		}
	}
	if isNone {
		fmt.Println("None")
	}
}
