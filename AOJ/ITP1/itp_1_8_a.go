package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	sc.Scan()
	s = sc.Text()
	for i := 0; i < len(s); i++ {
		s_upper := strings.ToUpper(string(s[i]))
		s_lower := strings.ToLower(string(s[i]))
		if string(s[i]) == s_upper {
			fmt.Print(s_lower)
		} else {
			fmt.Print(s_upper)
		}
	}
	fmt.Println()
}
