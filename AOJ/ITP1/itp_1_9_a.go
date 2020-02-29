package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	k := readString()
	var s string
	var cnt int
	for {
		s = readString()
		if s == "END_OF_TEXT" {
			break
		}
		s = strings.ToLower(s)
		if k == s {
			cnt++
		}
	}
	fmt.Println(cnt)
}
