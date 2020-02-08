package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)
	fmt.Printf("%s%s%s\n", strings.ToUpper(s1[0:1]), strings.ToUpper(s2[0:1]), strings.ToUpper(s3[0:1]))
}
