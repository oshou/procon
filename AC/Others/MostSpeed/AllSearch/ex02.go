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

func readInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	steps := map[string]int{
		"a": 2,
		"b": 1,
		"c": 1,
		"d": 2,
		"e": 1,
		"f": 2,
		"g": 1,
		"h": 1,
		"i": 2,
		"j": 2,
		"k": 3,
		"l": 1,
		"m": 1,
		"n": 1,
		"o": 1,
		"p": 1,
		"q": 1,
		"r": 1,
		"s": 1,
		"t": 2,
		"u": 1,
		"v": 1,
		"w": 1,
		"x": 2,
		"y": 2,
		"z": 1,
	}
	cnt := 0
	for _, v1 := range steps {
		for _, v2 := range steps {
			for _, v3 := range steps {
				if v1+v2+v3 <= 5 {
					cnt++
				}
			}
		}
	}
	fmt.Println("cnt:", cnt)
}
