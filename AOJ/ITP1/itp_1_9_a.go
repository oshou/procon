package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := readString()
	cnt := 0
	for {
		word := readString()
		if word == "END_OF_TEXT" {
			break
		}
		s = strings.ToLower(s)
		word = strings.ToLower(word)
		if word == s {
			cnt++
		}
	}
	fmt.Println(cnt)
}
