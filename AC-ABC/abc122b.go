package main

import "fmt"

func main() {
	var s string
	var max, tmp int = 0, 0
	fmt.Scan(&s)
	for _, c := range s {
		switch string(c) {
		case "A", "C", "G", "T":
			tmp++
		default:
			if max < tmp {
				max = tmp
				tmp = 0
			}
		}
	}
	if max < tmp {
		max = tmp
		tmp = 0
	}
	fmt.Println(max)
}
