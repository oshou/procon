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
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	n := readInt()
	fmt.Println("n:", n)
	nums := make([]bool, n+1)
	max := int(math.Floor(math.Sqrt(float64(n))))
	fmt.Println("max:", max)
	for i := 2; i <= max; i++ {
		for j := 2 * i; j <= n; j += i {
			nums[j] = true
		}
	}
	cnt := 0
	for i := 1; i <= n; i++ {
		if !nums[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
