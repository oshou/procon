package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a, b, c := readInt(), readInt(), readInt()
	nums := []int{a, b, c}
	sort.Ints(nums)
	fmt.Printf("%v %v %v\n", nums[0], nums[1], nums[2])
}
