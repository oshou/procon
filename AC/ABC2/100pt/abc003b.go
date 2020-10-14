package main

import "fmt"

func contain(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	var s, t string
	at := []string{"a", "t", "c", "o", "d", "e", "r"}
	fmt.Scan(&s)
	fmt.Scan(&t)
	ans := "You can win"
	for i := 0; i < len(s); i++ {
		chr_s := s[i : i+1]
		chr_t := t[i : i+1]
		if chr_s == chr_t {
			continue
		}
		if chr_s == "@" && contain(at, chr_t) {
			continue
		}
		if chr_t == "@" && contain(at, chr_s) {
			continue
		}
		ans = "You will lose"
	}
	fmt.Println(ans)
}
