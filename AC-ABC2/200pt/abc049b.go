package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const MaxBuf = 200100

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
	var h int
	h, _ = readInt(), readInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = readString()
		fmt.Println(s[i])
		fmt.Println(s[i])
	}
}
