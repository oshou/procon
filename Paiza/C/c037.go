package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := readString()
	fmt.Println(s)
	M, _ := strconv.Atoi(string(s[0:2]))
	d, _ := strconv.Atoi(string(s[3:5]))
	h, _ := strconv.Atoi(string(s[6:8]))
	m, _ := strconv.Atoi(string(s[9:11]))
	d += h / 24
	h = h % 24
	fmt.Printf("%02d/%02d %02d:%02d\n",M, d, h, m)
}
