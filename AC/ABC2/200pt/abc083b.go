package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func digits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	n, a, b := readInt(), readInt(), readInt()
	sum, digit := 0, 0
	for i := 1; i <= n; i++ {
		digit = digits(i)
		if a <= digit && digit <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
