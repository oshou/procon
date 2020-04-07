package main

import "fmt"

func main() {
	var n int
	var c string
	var h, le, ge float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c, &h)
		switch c {
		case "le":
			if le == 0 || h < le {
				le = h
			}
		case "ge":
			if ge == 0 || h > ge {
				ge = h
			}
		}
	}
	fmt.Printf("%.1f %.1f\n", ge, le)
}
