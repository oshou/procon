package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(alphabet); i++ {
		if c == string(alphabet[i]) {
			fmt.Println(string(alphabet[i+1]))
			break
		}
	}
}
