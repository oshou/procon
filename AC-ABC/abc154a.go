package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	s, t := readString(), readString()
	a, b := readInt(), readInt()
	u := readString()
	switch {
	case strings.Index(s, u) >= 0:
		a--
	case strings.Index(t, u) >= 0:
		b--
	}
	fmt.Println(a, b)
}
