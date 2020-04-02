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

func abs(x int) int {
	if x > 0 {
		return x
	}
	return (-1) * x
}

func main() {
	var sum, current, next int
	n := readInt()
	current = 1
	for i := 0; i < n; i++ {
		next = readInt()
		sum += abs(next - current)
		current = next
	}
	fmt.Println(sum)
}
