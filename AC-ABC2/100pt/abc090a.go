package main

import "fmt"

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)
	fmt.Printf("%c%c%c\n", s1[0], s2[1], s3[2])
}
