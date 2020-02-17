package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
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

func main() {
	h, w := readInt(), readInt()
	a := make([]string, h)
	for i := 0; i < h; i++ {
		a[i] = readString()
	}
	fmt.Println(strings.Repeat("#", w+2))
	for i := 0; i < h; i++ {
		fmt.Printf("#%v#\n", a[i])
	}
	fmt.Println(strings.Repeat("#", w+2))
}
