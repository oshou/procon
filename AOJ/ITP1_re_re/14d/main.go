package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n, min, max, sum int
		line             string
		sc               = bufio.NewScanner(os.Stdin)
	)

	fmt.Scan(&n)
	if sc.Scan() {
		line = sc.Text()
	}
	strs := strings.Split(line, " ")
	arr := make([]int, len(strs))
	for i, _ := range strs {
		arr[i], _ = strconv.Atoi(strs[i])
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			min = arr[i]
			max = arr[i]
			sum = arr[i]
			continue
		}

		if arr[i] < min {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
		sum += arr[i]
	}
	fmt.Println(min, max, sum)
}
