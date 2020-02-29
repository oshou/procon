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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func digits(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}
		sum += n % 10
		n = n / 10
	}
	return sum
}

func main() {
	var ans []int
	for {
		s := readInt()
		if s == 0 {
			break
		}
		ans = append(ans, digits(s))
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}
