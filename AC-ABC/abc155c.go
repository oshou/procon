package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	s := make(map[string]int)
	for i := 0; i < n; i++ {
		s[readString()]++
	}

	hack := map[int]string{}
	hackKeys := []int{}
	for k, v := range s {
		hack[v] = k
		hackKeys = append(hackKeys, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(hackKeys)))
	max := hackKeys[0]
	sort.Strings(s)
	for k, v := range s {
		if v == max {
			fmt.Println(k)
		}
	}
}
