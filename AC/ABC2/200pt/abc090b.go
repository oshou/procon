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
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a, b := readInt(), readInt()
	cnt := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		frac := true
		for j := 0; j <= len(s)/2-1; j++ {
			if s[j] != s[len(s)-j-1] {
				frac = false
			}
		}
		if frac == true {
			cnt++
		}
	}
	fmt.Println(cnt)
}
