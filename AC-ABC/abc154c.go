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
	a := make([]int, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		m[a[i]]++
	}
	if len(a) == len(m) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
