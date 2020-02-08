package main

import "fmt"

func main() {
	var c string
	fmt.Scanf("%s", &c)
	switch c {
	case "a", "i", "u", "e", "o":
		fmt.Println("vowel")
	default:
		fmt.Println("consonant")
	}
}
