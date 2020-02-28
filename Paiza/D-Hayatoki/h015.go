package main

import "fmt"

func main() {
	var t1, t2 int
	fmt.Scan(&t1, &t2)
	diff := t2 - t1
	switch {
	case diff > 0:
		fmt.Printf("+%d\n", diff)
	default:
		fmt.Printf("%d\n", diff)
	}
}
