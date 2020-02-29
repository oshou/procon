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
	var n, x int
	var k, cnt int
	for {
		k, cnt = 0, 0
		n, x = readInt(), readInt()
		if n == 0 && k == 0 {
			os.Exit(0)
		}
		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				k = x - (i + j)
				if j+1 <= k && k <= n && i+j+k == x {
					cnt++
				}
			}
		}
		fmt.Println(cnt)
	}
}
