package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	arr := make([]string, 3)
	arr = strings.Split(sc.Text(), " ")
	for _, v := range arr {
		fmt.Println(v)
	}
}
