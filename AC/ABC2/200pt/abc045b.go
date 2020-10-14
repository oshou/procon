package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	arr := map[string]string{"a": a, "b": b, "c": c}
	key := "a"
	for arr[key] != "" {
		c := string(arr[key][0])
		arr[key] = arr[key][1:]
		key = c
	}
	fmt.Println(strings.ToUpper(key))
}
