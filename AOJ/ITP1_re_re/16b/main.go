package main

import "fmt"

var (
	symbol string
	number int
)

func main() {
	spade := make([]bool, 14)
	hart := make([]bool, 14)
	club := make([]bool, 14)
	daiya := make([]bool, 14)
	for i := 0; i <= 13; i++ {
		spade[i] = false
		hart[i] = false
		club[i] = false
		daiya[i] = false
	}

	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&symbol, &number)
		switch symbol {
		case "S":
			spade[number] = true
		case "H":
			hart[number] = true
		case "C":
			club[number] = true
		case "D":
			daiya[number] = true
		}
	}

	for i := 1; i <= 13; i++ {
		if !spade[i] {
			fmt.Println("S", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !hart[i] {
			fmt.Println("H", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !club[i] {
			fmt.Println("C", i)
		}
	}
	for i := 1; i <= 13; i++ {
		if !daiya[i] {
			fmt.Println("D", i)
		}
	}
}
