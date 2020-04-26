package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sn string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()
	for i := 0; i < len(s); i++ {
		switch {
		case 'a' <= s[i] && s[i] <= 'z':
			sn += strings.ToUpper(string(s[i]))
		case 'A' <= s[i] && s[i] <= 'Z':
			sn += strings.ToLower(string(s[i]))
		default:
			sn += string(s[i])
		}
	}
	fmt.Println(sn)
}
