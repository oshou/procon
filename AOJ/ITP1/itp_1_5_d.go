package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	n := readInt()
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Printf(" %v", i)
			continue
		}
		n := i
		for {
			if n == 0 {
				break
			}
			if n%10 == 3 {
				fmt.Printf(" %v", i)
				break
			}
			n /= 10
		}
	}
	fmt.Println()
}
