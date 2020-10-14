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
	for i := 0; i < n; i++ {
		arr = append(arr, readInt())
	}
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			if arr[i]%3 != 0 && arr[i]%5 != 0 {
				fmt.Println("DENIED")
				os.Exit(0)
			}
		}
	}
	fmt.Println("APPROVED")
}
