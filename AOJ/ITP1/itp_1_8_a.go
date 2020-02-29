package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := readString()
	var ans string = ""
	for _, c := range s {
		switch {
		case 'a' <= c && c <= 'z':
			ans += strings.ToLower(string(c))
		case 'A' <= c && c <= 'Z':
			ans += strings.ToLower(string(c))
		default:
			ans += string(c)
		}
	}
	fmt.Println(ans)
}
