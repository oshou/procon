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
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	var h, w int
	for {
		h, w = readInt(), readInt()
		if h == 0 && w == 0 {
			os.Exit(0)
		}
		fmt.Println(strings.Repeat("#", w))
		for i := 0; i < h-2; i++ {
			fmt.Printf("#%v#\n", strings.Repeat(".", w-2))
		}
		fmt.Println(strings.Repeat("#", w))
		fmt.Println("")
	}
}
