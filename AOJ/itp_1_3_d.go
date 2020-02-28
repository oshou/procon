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
	a, b, c := readInt(), readInt(), readInt()
	cnt := 0
	for i := a; i <= b; i++ {
		switch {
		case i == 0:
			continue
		case c%i == 0:
			cnt++
		}
	}
	fmt.Println(cnt)
}
