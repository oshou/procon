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
	var n, tmp int
	n = readInt()
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Printf(" %v", i)
			continue
		}
		tmp = i
		for tmp > 0 {
			if tmp%10 == 3 {
				fmt.Printf(" %v", i)
				break
			}
			tmp /= 10
		}
	}
	fmt.Println()
}
