package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
