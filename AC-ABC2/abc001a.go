package main

import "fmt"

func main() {
	var m int
	var km float64
	fmt.Scan(&m)
	km = m / 1000
	switch {
	case km < 0.1:
		fmt.Println("00")
	case 0.1 <= km && km < 1:
		fmt.Printf("0%d\n", 10*km)
	case 1 <= km && km <= 5:
		fmt.Println(10 * km)
	case 6 <= km && km <= 30:
		fmt.Println(km + 50)
	case 35 <= km && km <= 70:
		fmt.Println((km-30)/5 + 80)
	case 70 < km:
		fmt.Println("89")
	}
}
