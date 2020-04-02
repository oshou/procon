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
	n, x := readInt(), readInt()
	var k int
	var sbstr string = fmt.Sprintf("%b", x)
	for i := 1; i <= n; i++ {
		k = readInt()
		fmt.Println(string(sbstr[len(sbstr)-k]))
	}
}
