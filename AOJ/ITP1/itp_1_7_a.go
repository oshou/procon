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
	var m, f, r, score int
	for {
		m, f, r = readInt(), readInt(), readInt()
		score = m + f
		if m == -1 && f == -1 && r == -1 {
			os.Exit(0)
		}
		switch {
		case m == -1 || f == -1:
			fmt.Println("F")
		case 80 <= score:
			fmt.Println("A")
		case 65 <= score && score < 80:
			fmt.Println("B")
		case (50 <= score && score < 65) || (50 <= r):
			fmt.Println("C")
		case 30 <= score && score < 50:
			fmt.Println("D")
		case score < 30:
			fmt.Println("F")
		}
	}
}
