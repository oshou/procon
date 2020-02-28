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
	n := readInt()
	for i := 0; i < n; i++ {
		t, x, y := readInt(), readInt(), readInt()
		if x+y > t || (x+y+t)%2 != 0 {
			fmt.Println("No")
			os.Exit(0)
		}
	}
	fmt.Println("Yes")
}
