package main

import "fmt"

func main() {
	nums := make([]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&nums[i])
	}
	var yes, no int
	for _, num := range nums {
		if num == "yes" {
			yes++
		} else {
			no++
		}
	}
	if yes > no {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
