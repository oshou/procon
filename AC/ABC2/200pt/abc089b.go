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
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m[readString()]++
	}
	len := len(m)
	switch len {
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	}
}
