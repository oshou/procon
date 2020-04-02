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
	x, p := readInt(), readInt()
	var sum int = x
	for x != 0 {
		x = int(x * (100 - p) / 100)
		sum += x
	}
	fmt.Println(sum)
}
