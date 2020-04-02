package main

import "fmt"

var alphabets = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func next(s string, n int) string {
	for i := range alphabets {
		if alphabets[i] == s {
			if len(alphabets)-i >= n {
				return alphabets[i+n]
			} else {
				return alphabets[(n-len(alphabets)+i)%n]
			}
		}
	}
	return ""
}

func back(s string, n int) string {
	for i := range alphabets {
		if alphabets[i] == s {
			if i > n {
				return alphabets[i-n]
			} else {
				return alphabets[(i-n+len(alphabets))%n]
			}
		}
	}
	return ""
}

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			fmt.Print(back(string(s[i]), n))
		} else {
			fmt.Print(next(string(s[i]), n))
		}
	}
	fmt.Println()
}
