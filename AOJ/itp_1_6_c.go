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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := readInt()
	var p [4][3][10]int
	for i := 0; i < n; i++ {
		b, f, r, v := readInt(), readInt(), readInt(), readInt()
		p[b-1][f-1][r-1] = max(0, p[b-1][f-1][r-1]+v)
	}
	for i := 0; i < 4; i++ {
		if i != 0 {
			fmt.Println("####################")
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %v", p[i][j][k])
			}
			fmt.Println("")
		}
	}
}
