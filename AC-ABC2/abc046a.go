package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contain(nums []int, n int) bool {
	for _, num := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func main() {
	// input
	//var a, b, c int
	//fmt.Scanf("%d %d %d", &a, &b, &c)
	//nums := []int{a, b, c}
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nums := strings.Split(sc.Text(), " ")

	// create unique-num-slice
	uniques := []int{}
	for _, num := range nums {
		if !contain(uniques, num) {
			uniques = append(uniques, num)
		}
	}

	// count
	fmt.Println(len(uniques))
}
