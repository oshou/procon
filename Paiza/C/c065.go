package main

import "fmt"

func main() {
	var n, x, min, max int
	var isCollect bool
	var o string
	var div []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&o, &x)
		switch o {
		case ">":
			if min == 0 || min < x {
				min = x + 1
			}
		case "<":
			if max == 0 || x < max {
				max = x - 1
			}
		case "/":
			div = append(div, x)
		}
	}
	for i := min; i <= max; i++ {
		isCollect = true
		for _, v := range div {
			if i%v != 0 {
				isCollect = false
			}
		}
		if isCollect {
			fmt.Println(i)
			break
		}
	}
}
