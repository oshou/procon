package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var c []string
	var sc = bufio.NewScanner(os.Stdin)
	c = make([]string, n)

	fmt.Scan(&n)
	sc.Scan()
	c = strings.Split(sc.Text(), " ")

	var hasZero bool
	var hasX10 bool
	var max int = 0
	var sum int = 0
	for i := 0; i < n; i++ {
		switch c[i] {
		case "x10":
			hasX10 = true
		case "0":
			hasZero = true
		default:
			num, _ := strconv.Atoi(c[i])
			sum += num
			if max < num {
				max = num
			}
		}
	}
	if hasZero {
		sum -= max
	}
	if hasX10 {
		sum *= 10
	}
	fmt.Println(sum)
}
