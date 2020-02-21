package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	n := readInt()
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, readInt())
	}
	max, min, sum := -1000000, 1000000, 0
	for _, v := range arr {
		if max < v {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	fmt.Printf("%v %v %v\n", min, max, sum)
}
