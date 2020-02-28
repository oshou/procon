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
	s := readString()
	max := 0
	for i := 0; i < n; i++ {
		left := make(map[string]bool)
		right := make(map[string]bool)
		for j := 0; j < n; j++ {
			if j <= i {
				left[string(s[j])] = true
			} else {
				right[string(s[j])] = true
			}
		}
		cnt := 0
		for k := range left {
			if left[k] == right[k] {
				cnt++
			}
		}
		if max < cnt {
			max = cnt
		}
	}
	fmt.Println(max)
}
