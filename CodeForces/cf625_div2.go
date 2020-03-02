package main

import (
	"bufio"
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
	n := readInt()
	r := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = readInt()
	}
	for i := 0; i < n; i++ {
		b[i] = readInt()
	}
	win, lose := 0, 0
	for i := 0; i < n; i++ {
		switch {
		case r[i] == 1 && b[i] == 0:
			win++
		case r[i] == 0 && b[i] == 1:
			lose++
		}
	}
	ans := 0
	switch {
	case win == 0:
		ans = "-1"
	case lose == 0:
		ans = "1"
	default:
		ans = win/lose + 1
	}
	fmt.Println(ans)
}
