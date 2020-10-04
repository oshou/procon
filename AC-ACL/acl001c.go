package main

import "fmt"

type unionFind struct {
	count   int
	parents []int
}

func NewUnionFind(n int) *unionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &unionFind{
		count:   n,
		parents: p,
	}
}

func (u *unionFind) Root(x int) int {
	if u.parents[x] == x {
		return x
	}
	u.parents[x] = u.Root(u.parents[x])
	return u.parents[x]
}

func (u *unionFind) Unite(x, y int) {
	rx, ry := u.Root(x), u.Root(y)
	switch {
	case rx == ry:
		return
	case rx < ry:
		u.parents[rx] = ry
	case rx > ry:
		u.parents[ry] = rx
	}
}

func (u *unionFind) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)
	var cnt = n
	u := NewUnionFind(n)
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		if u.Same(a-1, b-1) {
			continue
		} else {
			u.Unite(a-1, b-1)
			cnt--
		}
	}
	if cnt > 0 {
		fmt.Println(cnt - 1)
	} else {
		fmt.Println(cnt)
	}
}
