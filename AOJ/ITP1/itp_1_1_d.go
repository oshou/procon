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
	S := readInt()
	h := S / 3600
	m := (S % 3600) / 60
	s := (S % 3600) % 60
	fmt.Printf("%v:%v:%v\n", h, m, s)
}
