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
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	if len(m) == len(s) {
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
}
