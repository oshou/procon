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
	w, h, x, y, r := readInt(), readInt(), readInt(), readInt(), readInt()
	if (0 <= x-r) && (x+r <= w) && (0 <= y-r) && (y+r <= h) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
