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
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	var sum int
	n, m, l := readInt(), readInt(), readInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			a[i][j] = readInt()
		}
	}
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		for j := 0; j < l; j++ {
			b[i][j] = readInt()
		}
	}
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(c[i][j])
		}
		fmt.Println()
	}
}
