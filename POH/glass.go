package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// n
	var n int
	fmt.Scan(&n)
	// input
	input := make([][]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		input[i] = strings.Split(scanner.Text(), " ")
	}

	// m
	var m int
	fmt.Scan(&m)
	// pattern
	pattern := make([][]string, m)
	for j := 0; j < m; j++ {
		scanner.Scan()
		pattern[j] = strings.Split(scanner.Text(), " ")
	}

	for i := 0; i <= n-m; i++ {
		for j := 0; j <= n-m; j++ {
		}
	}
}
