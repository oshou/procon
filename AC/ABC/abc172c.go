// 最短経路問題?
// 貪欲法はだめそう。目先のa,b比較だとエッジケースに対応できない
// 長い目でみて最適かどうかをどうチェックするか
package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func maketree(parent *node, newValue int) {
	var child node
	child.value =
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	q := make([]int, 0)
}
