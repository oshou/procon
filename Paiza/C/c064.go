package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	c := make([]int, m)
	var sum int
	var a []string
	for i := 0; i < m; i++ {
		fmt.Scan(&c[i])
	}
	for i := 0; i < n; i++ {
		a = strings.Split(readString(), " ")
		sum = 0
		for j := 0; j < m; j++ {
			aj, _ := strconv.ParseFloat(a[j], 0)
			sum += int(math.Floor(aj * float64(c[j]) / 100))
		}
		fmt.Println(sum)
	}
}
