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
	ans := ""
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		switch c {
		case "a", "i", "u", "e", "o":
			continue
		default:
			ans += c
		}
	}
	fmt.Println(ans)
}
