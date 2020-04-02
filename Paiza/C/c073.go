package main

import (
	"fmt"
	"math"
)

func main() {
	var L, u, a, b, v float64
	fmt.Scan(&L)
	fmt.Scan(&u, &a, &b)
	fmt.Scan(&v)
	rabbit := L*u + (math.Floor(L/a)-1)*b
	turtle := L * v
	switch {
	case rabbit == turtle:
		fmt.Println("DRAW")
	case rabbit < turtle:
		fmt.Println("USAGI")
	default:
		fmt.Println("KAME")
	}
}
