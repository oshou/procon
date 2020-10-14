package main

import "fmt"

func main() {
	// input
	var w, h, n int
	fmt.Scan(&w, &h, &n)
	x := make([]int, n)
	y := make([]int, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i], &a[i])
	}

	// set x,y max&min
	var xmin int = 0
	var xmax int = w
	var ymin int = 0
	var ymax int = h
	for i := 0; i < n; i++ {
		switch a[i] {
		case 1:
			if xmin < x[i] {
				xmin = x[i]
			}
		case 2:
			if x[i] < xmax {
				xmax = x[i]
			}
		case 3:
			if ymin < y[i] {
				ymin = y[i]
			}
		case 4:
			if y[i] < ymax {
				ymax = y[i]
			}
		}
	}

	// calculate area
	if xmax <= xmin || ymax <= ymin {
		fmt.Println(0)
	} else {
		fmt.Println((xmax - xmin) * (ymax - ymin))
	}
}
