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

func main() {
	var t string
	var e, m, s, j, g, total, major, cnt int
	n := readInt()
	for i := 0; i < n; i++ {
		t, e, m, s, j, g = readString(), readInt(), readInt(), readInt(), readInt(), readInt()
		total = e + m + s + j + g
		switch t {
		case "s":
			major = m + s
		case "l":
			major = j + g
		}
		if total >= 350 && major >= 160 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
