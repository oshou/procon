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
	s := readString()
	s = strings.Replace(s, "eraser", "", -1)
	s = strings.Replace(s, "erase", "", -1)
	s = strings.Replace(s, "dreamer", "", -1)
	s = strings.Replace(s, "dream", "", -1)
	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
