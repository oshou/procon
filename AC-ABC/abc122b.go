package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var max, tmp int = 0, 0
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
