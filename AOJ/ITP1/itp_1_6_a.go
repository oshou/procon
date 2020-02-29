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
	n := readInt()
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, readInt())
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			fmt.Printf("%v", arr[i])
		} else {
			fmt.Printf(" %v", arr[i])
		}
	}
	fmt.Println("")
}
