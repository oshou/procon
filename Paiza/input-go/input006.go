package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]string, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
