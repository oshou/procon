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

type calc struct {
	a, b int
	op   string
}

func main() {
	var a, b int
	var op string
	var calcs []calc
	var ans int
	for {
		a, op, b = readInt(), readString(), readInt()
		if op == "?" {
			break
		}
		calcs = append(calcs, calc{a: a, b: b, op: op})
	}
	for _, calc := range calcs {
		switch calc.op {
		case "+":
			ans = calc.a + calc.b
		case "-":
			ans = calc.a - calc.b
		case "*":
			ans = calc.a * calc.b
		case "/":
			ans = calc.a / calc.b
		}
		fmt.Println(ans)
	}
}
