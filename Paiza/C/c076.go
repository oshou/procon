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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	x, y, z := readInt(), readInt(), readInt()
	n := readInt()
	s := make([]int, n)
	t := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		s[i], t[i] = readInt(), readInt()

		if s[i] <= 9 && 0 <= t[i] {
			tmp := z * (min(9, t[i]) - max(0, s[i]))
			sum += tmp
		}

		if s[i] <= 17 && 9 <= t[i] {
			tmp := x * (min(17, t[i]) - max(9, s[i]))
			sum += tmp
		}

		if s[i] <= 22 && 17 <= t[i] {
			tmp := y * (min(22, t[i]) - max(17, s[i]))
			sum += tmp
		}

		if s[i] <= 24 && 22 <= t[i] {
			tmp := z * (min(24, t[i]) - max(22, s[i]))
			sum += tmp
		}
	}
	fmt.Println(sum)
}
