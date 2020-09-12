package main

func reverse(x int) int {
	var num, reversed int
	var isNegative bool = false
	if x < 0 {
		isNegative = true
		num = -x
	} else {
		num = x
	}
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	if isNegative {
		return (-1) * reversed
	}
	return reversed
}
