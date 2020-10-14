package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var arr []int
	a, b := 0, 0
	for i := 0; i < n; i++ {
		arr = append(arr, readInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a += arr[i]
		} else {
			b += arr[i]
		}
	}
	fmt.Println(a - b)
}
