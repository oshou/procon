package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	t1, _ := time.Parse("2006/01/02", s)
	t2, _ := time.Parse("2006/01/02", "2019/04/30")
	if t1.After(t2) {
		fmt.Println("TBD")
	} else {
		fmt.Println("Heisei")
	}
}
