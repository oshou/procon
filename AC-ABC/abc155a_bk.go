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
	sc.Buffer([]byte{}, math.Maxint64)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a, b, c := readInt(), readInt(), readInt()
	if a == b && b == c {
		fmt.Println("No")
	}
	if a == b || b == c || c == a {
		fmt.Println("Yes")
	}
	fmt.Println("No")
}
