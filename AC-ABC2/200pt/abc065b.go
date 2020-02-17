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
	n := readInt()
	a := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		a[i] = readInt()
	}
	cnt := 1
	i := 1
	for {
		switch {
		case a[i] == 2:
			fmt.Println(cnt)
			os.Exit(0)
		case a[i] == -1:
			fmt.Println(-1)
			os.Exit(0)
		default:
			cnt++
			i, a[i] = a[i], -1
		}
	}
}
