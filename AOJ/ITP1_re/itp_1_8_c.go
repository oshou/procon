package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int)
	sc := bufio.NewReader(os.Stdin)
	for {
		s, err := sc.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, cbyte := range s {
			cchar := strings.ToLower(string(cbyte))
			if "a" <= cchar && cchar <= "z" {
				m[cchar]++
			}
		}
	}
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, key := range keys {
		fmt.Printf("%s : %d\n", key, m[key])
	}
}
