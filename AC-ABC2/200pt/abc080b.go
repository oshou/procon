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

func main() {
	s := readString()
	x, _ := strconv.Atoi(s)
	fx := 0
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		fx += v
	}
	if x%fx == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
