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
	n, y := readInt(), readInt()
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if 10000*i+5000*j+1000*(n-i-j) == y {
				fmt.Printf("%v %v %v\n", i, j, n-i-j)
				os.Exit(0)
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
