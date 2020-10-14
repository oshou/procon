package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	o, e := readString(), readString()
	pass := ""
	for i := 0; i < len(e); i++ {
		pass += string(o[i])
		pass += string(e[i])
	}
	if len(o) != len(e) {
		pass += string(o[len(o)-1])
	}
	fmt.Println(pass)
}
