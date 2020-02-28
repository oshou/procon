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
	sc.Split(bufio.ScanWords())
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
	a, b := readInt(), readInt()
	s := readString()
	if len(arr[0]) != a || len(arr[1] != b) {
		fmt.Println("No")
	}
	c := ""
	for i := 0; i < len(s); i++ {
		c = string(s[i])
		switch {
		case i < a:
		case a < i:
			if !("0" <= c && c <= "9") {
				fmt.Println("No")
				break
			}
		default:
			if c != "-" {
				fmt.Println("No")
				break
			}
		}
	}
	fmt.Println("Yes")
}
