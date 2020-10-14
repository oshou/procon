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

func reverse(s string) string {
	var ans string
	for i := len(s) - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	return ans
}

func main() {
	var s string
	var q int
	var o, f int
	var c string
	s = readString()
	q = readInt()
	for i := 1; i <= q; i++ {
		o = readInt()
		switch o {
		case 1:
			s = reverse(s)
		case 2:
			f, c = readInt(), readString()
			switch f {
			case 1:
				s = c + s
			case 2:
				s = s + c
			}
		}
	}
	fmt.Println(s)
}
