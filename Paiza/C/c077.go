package main

import "fmt"

func main() {
	var k, n int
	var p float64
	var grade string
	fmt.Scan(&k, &n)
	d := make([]int, k)
	a := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&d[i], &a[i])
		switch {
		case d[i] >= 10:
			fmt.Println("D")
			continue
		case d[i] > 0:
			p = (float64(a[i]) / float64(n)) * 100 * 0.8
		default:
			p = (float64(a[i]) / float64(n)) * 100
		}
		switch {
		case 80 <= p:
			grade = "A"
		case 70 <= p:
			grade = "B"
		case 60 <= p:
			grade = "C"
		default:
			grade = "D"
		}
		fmt.Println(grade)
	}
}
