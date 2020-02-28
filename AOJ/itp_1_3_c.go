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
	x, y := 0, 0
	arrx, arry := []int{}, []int{}
	for {
		x, y = readInt(), readInt()
		if x == 0 && y == 0 {
			break
		}
		arrx = append(arrx, x)
		arry = append(arry, y)
	}
	for i := 0; i < len(arrx); i++ {
		if arrx[i] < arry[i] {
			fmt.Printf("%v %v\n", arrx[i], arry[i])
		} else {
			fmt.Printf("%v %v\n", arry[i], arrx[i])
		}
	}
}
