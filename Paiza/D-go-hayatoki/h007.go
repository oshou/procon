package main

import "fmt"

func main() {
	boldLevel := []string{"6B", "5B", "4B", "3B", "2B", "B", "HB", "F", "H", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H"}
	var k int
	fmt.Scan(&k)
	fmt.Println(boldLevel[k-1])
}
