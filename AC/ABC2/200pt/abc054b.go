package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	// Input
	n, m := readInt(), readInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString()
	}
	b := make([]string, m)
	for i := 0; i < m; i++ {
		b[i] = readString()
	}

	l := n - m
	for i := 0; i <= l; i++ {
		for j := 0; j <= l; j++ {
			tmp := make([]string, m)
			for k := 0; k < m; k++ {
				tmp[k] = a[k+j][i : m+i]
			}
			if reflect.DeepEqual(tmp, b) {
				fmt.Println("Yes")
				os.Exit(0)
			}
		}
	}
	fmt.Println("No")
}
