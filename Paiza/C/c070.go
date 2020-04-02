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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func mapmax(m map[string]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var s, ans string
	var m map[string]int
	n := readInt()
	for i := 0; i < n; i++ {
		s = readString()
		m = make(map[string]int)
		for j := 0; j < len(s); j++ {
			m[string(s[j])]++
		}
		switch {
		case mapmax(m) == 4:
			ans = "Four Card"
		case mapmax(m) == 3:
			ans = "Three Card"
		case mapmax(m) == 2 && len(m) == 2:
			ans = "Two Pair"
		case mapmax(m) == 2 && len(m) == 3:
			ans = "One Pair"
		case len(m) == 4:
			ans = "No Pair"
		}
		fmt.Println(ans)
	}
}
