package main

import "fmt"

func main() {
	var ans string
	var d int
	fmt.Scan(&d)
	switch d {
	case 25:
		ans = "Christmas"
	case 24:
		ans = "Christmas Eve"
	case 23:
		ans = "Christmas Eve Eve"
	case 22:
		ans = "Christmas Eve Eve Eve"
	}
	fmt.Println(ans)
}
