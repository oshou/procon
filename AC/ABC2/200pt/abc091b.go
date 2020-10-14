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
	sm := make(map[string]int)
	for i := 0; i < n; i++ {
		sm[readString()]++
	}
	var keys []string
	for k := range sm {
		keys = append(keys, k)
	}

	m := readInt()
	tm := make(map[string]int)
	for i := 0; i < m; i++ {
		tm[readString()]++
	}
	max := 0
	for _, v := range keys {
		v := sm[v] - tm[v]
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
