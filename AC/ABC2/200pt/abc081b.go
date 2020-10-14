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

func div(arr []int) ([]int, bool) {
	div := []int{}
	for _, v := range arr {
		if v%2 != 0 {
			return arr, false
		}
		div = append(div, v/2)
	}
	return div, true
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	cnt := 0
	for {
		if div, ok := div(arr); ok {
			cnt++
			arr = div
			continue
		}
		break
	}
	fmt.Println(cnt)
}
