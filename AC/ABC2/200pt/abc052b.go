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
	cnt, max := 0, 0
	for i := 0; i < n; i++ {
		switch string(s[i]) {
		case "I":
			cnt++
			if max < cnt {
				max = cnt
			}
		case "D":
			cnt--
		}
	}
	fmt.Println(max)
}
