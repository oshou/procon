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

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a, b, c := readInt(), readInt(), readInt()
	switch {
	case a == b && b == c:
		fmt.Println("No")
	case a == b || b == c || c == a:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}
}
