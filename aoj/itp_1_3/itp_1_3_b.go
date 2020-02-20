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
	var nums []int
	for {
		x := readInt()
		if x == 0 {
			break
		}
		nums = append(nums, x)
	}
	for i, num := range nums {
		fmt.Printf("Case %v: %v\n", i+1, num)
	}
}
