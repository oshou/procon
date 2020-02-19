
package main

import (
	"bufio"
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
	n := readString()
	for i:= 0;i<len(s);i++{
		fmt.Println(i)
	}
}
