package main

import (
	"fmt"
	"time"
)

type Loc struct {
	loc  string
	diff int
}

func main() {
	var n, s int
	var p string
	var q, t string
	fmt.Scan(&n)
	l := make(map[string]int)
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p, &s)
		keys[i] = p
		l[p] = s
	}
	fmt.Scan(&q, &t)
	layout := "15:04"
	tt, _ := time.Parse(layout, t)
	for _, v := range keys {
		fmt.Println(tt.Add(time.Duration(l[v]-l[q]) * time.Hour).Format(layout))
	}
}
