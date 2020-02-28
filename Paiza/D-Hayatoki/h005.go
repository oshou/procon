package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
	fmt.Println(strings.Repeat("^", len(scanner.Text())))
}
