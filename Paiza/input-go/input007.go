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
	arr := strings.Split(sc.Text(), ",")
	for _, ele := range arr {
		fmt.Println(ele)
	}
}
