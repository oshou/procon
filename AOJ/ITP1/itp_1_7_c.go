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

func readInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	var r, c int = readInt(), readInt()
	arr := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		arr[i] = make([]int, c+1)
	}

	for i := 0; i < r; i++ {
		total := 0
		for j := 0; j < c; j++ {
			arr[i][j] = readInt()
			total += arr[i][j]
			arr[r][j] += arr[i][j]
		}
		arr[i][c] = total
		arr[r][c] += total
	}

	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}
