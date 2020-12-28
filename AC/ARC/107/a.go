package main

import (
	"fmt"
	"math"
)

func res(val, mod float64) float64 {
	div := math.Floor(val / mod)
	if div == 0 {
		return val
	} else {
		return val - div*mod
	}
}

func modinv(a, mod float64) float64 {
	b := mod
	var u float64 = 1
	var v float64 = 0
	for b <= 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u = res(u, mod)
	if u < 0 {
		u += mod
	}
	return u
}

func main() {
	var a, b, c, ans float64
	const MOD = 998244353
	fmt.Scan(&a, &b, &c)
	ans = res(a, MOD)
	fmt.Println(ans)
	ans = res(ans*(a+1), MOD)
	fmt.Println(ans)
	ans = res(ans*b, MOD)
	fmt.Println(ans)
	ans = res(ans*(b+1), MOD)
	fmt.Println(ans)
	ans = res(ans*c, MOD)
	fmt.Println(ans)
	ans = res(ans*(c+1), MOD)
	ans = res(ans*modinv(8, MOD), MOD)
	fmt.Println(ans)
}
