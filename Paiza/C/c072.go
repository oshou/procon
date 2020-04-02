package main

import "fmt"

func main() {
	var n, atk, def, agi int
	var canEvolve bool
	fmt.Scan(&atk, &def, &agi)
	fmt.Scan(&n)
	s := make([]string, n)
	minatk := make([]int, n)
	maxatk := make([]int, n)
	mindef := make([]int, n)
	maxdef := make([]int, n)
	minagi := make([]int, n)
	maxagi := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &minatk[i], &maxatk[i], &mindef[i], &maxdef[i], &minagi[i], &maxagi[i])
		if minatk[i] <= atk && atk <= maxatk[i] && mindef[i] <= def && def <= maxdef[i] && minagi[i] <= agi && agi <= maxagi[i] {
			canEvolve = true
			fmt.Println(s[i])
		}
	}
	if !canEvolve {
		fmt.Println("no evolution")
	}
}
