package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	var cnt int
	sc.Split(bufio.ScanWords)
	n, d := nextInt(), nextInt()
	x := make([]int, n)
	y := make([]int, n)
	d2 := int(d * d)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
		if d2 >= x[i]*x[i]+y[i]*y[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}

//package b
//
//func numPoints(n, d int, x, y []int) int {
//	// d^2 = i^2+j^2で一致確認
//	// これを0からn-1まで実施する
//	// O(n)でn=2x10^5なので1x10^8である上限はクリアできる見込み
//	var cnt int
//	d2 := d * d
//	for i := 0; i < n; i++ {
//		if d2 >= x[i]*x[i]+y[i]*y[i] {
//			cnt += 1
//		}
//	}
//	return cnt
//}
