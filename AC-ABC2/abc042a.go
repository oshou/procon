package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// input
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")

	// count
	cnt7 := 0
	cnt5 := 0
	for i := 0; i < 3; i++ {
		if arr[i] == "7" {
			cnt7++
		}
		if arr[i] == "5" {
			cnt5++
		}
	}

	// ans
	if cnt7 == 1 && cnt5 == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
