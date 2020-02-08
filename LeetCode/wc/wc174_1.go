package main

import "fmt"

func count(nums []int, x int) int {
	cnt := 0
	for _, num := range nums {
		if num == x {
			cnt++
		}
	}
	return cnt
}

func kWeakestRows(mat [][]int, k int) []int {
	m := make(map[int]int)
	for i, row := range mat {
		fmt.Println(row)
		m[i] = count(row, 1)
	}
	values := make([]int, n)
	for _, v := range m {
		append
	}
	return m
}

func main() {
	mat := [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}
	k := 2
	fmt.Println(kWeakestRows(mat, k))
}
