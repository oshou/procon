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
		switch {
		case "a" <= c && c <= "z":
			ans += strings.ToUpper(c)
		case "A" <= c && c <= "Z":
			ans += strings.ToLower(c)
		default:
			ans += c
		}
	}
	fmt.Println(ans)
}
