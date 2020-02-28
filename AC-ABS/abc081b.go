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
	var arr []int
	var cnt int
	for i := 0; i < n; i++ {
		arr = append(arr, readInt())
	}
	for {
		div := true
		for i := 0; i < n; i++ {
			if arr[i]%2 != 0 {
				div = false
				break
			}
			arr[i] = arr[i] / 2
		}
		if div == false {
			break
		}
		cnt++
	}
	fmt.Println(cnt)
}
