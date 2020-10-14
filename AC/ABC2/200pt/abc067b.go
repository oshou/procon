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
	n, k := readInt(), readInt()
	var l []int
	for i := 0; i < n; i++ {
		l = append(l, readInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	sum := 0
	for i := 0; i < k; i++ {
		sum += l[i]
	}
	fmt.Println(sum)
}
