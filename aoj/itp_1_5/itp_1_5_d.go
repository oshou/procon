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
	n := readInt()
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%10 == 3 {
			fmt.Print(" ")
			fmt.Print(i)
		}
	}
	fmt.Println()
}
