package main

import (
	"fmt"
	"math/big"
)

const LIMIT = 1000000000000000000

func main() {
	var n int
	fmt.Scan(&n)
	sum := big.NewInt(1)
	tmp := new(big.Int)
	for i := 0; i < n; i++ {
		fmt.Scan(tmp)
		sum = big.NewInt(0).Mul(sum, tmp)
		fmt.Println("sum:", sum.String())
		if sum.Cmp(big.NewInt(LIMIT)) == 1 {
			fmt.Println("-1")
			return
		}
	}
	fmt.Println(sum.String())
}
