package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	a := new(big.Float)
	b := new(big.Float)
	fmt.Scan(a, b)
	mul := big.NewFloat(0).Mul(a, b)
	fmt.Println(strings.Split(mul.String(), ".")[0])
}
