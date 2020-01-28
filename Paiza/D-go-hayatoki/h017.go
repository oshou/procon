package main

import "fmt"

func main() {
	var I int
	fmt.Scan(&I)
	switch {
	case I < 30:
		fmt.Println("quiet")
	case 30 <= I && I < 50:
		fmt.Println("normal")
	case 50 <= I && I < 70:
		fmt.Println("noisy")
	case 70 <= I:
		fmt.Println("very noisy")
	}
}
