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

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a, b := readInt(), readInt()
	var sum, digit int
	for a != 0 || b != 0 {
		sum += ((a%10 + b%10) % 10) * int(math.Pow(10, float64(digit)))
		a = a / 10
		b = b / 10
		digit++
	}
	zerofill := digit - len(strconv.Itoa(sum))
	fmt.Printf("%v%v\n", strings.Repeat("0", zerofill), sum)
}
