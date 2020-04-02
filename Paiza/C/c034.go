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
	var a, b, c, op string
	var ai, bi, ci int
	var ans int
	a, op, b, _, c = readString(), readString(), readString(), readString(), readString()
	switch {
	case c == "x":
		ai, _ = strconv.Atoi(a)
		bi, _ = strconv.Atoi(b)
		switch op {
		case "+":
			ans = ai + bi
		case "-":
			ans = ai - bi
		}
	case a == "x":
		bi, _ = strconv.Atoi(b)
		ci, _ = strconv.Atoi(c)
		switch op {
		case "+":
			ans = ci - bi
		case "-":
			ans = ci + bi
		}
	case b == "x":
		ai, _ = strconv.Atoi(a)
		ci, _ = strconv.Atoi(c)
		switch op {
		case "+":
			ans = ci - ai
		case "-":
			ans = ai - ci
		}
	}
	fmt.Println(ans)
}
