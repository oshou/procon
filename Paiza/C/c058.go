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

func roll(s string) string {
	return string(s[1:]) + string(s[0])
}

func main() {
	_, t, s := readInt(), readString(), readString()
	var cnt int
	for t != s {
		s = roll(s)
		cnt++
	}
	fmt.Println(cnt)
}
