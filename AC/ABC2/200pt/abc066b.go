package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	s := readString()
	for i := len(s) - 1; i >= 0; i-- {
		if i%2 == 0 && (s[:i/2]+s[:i/2] == s[:i]) {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}
