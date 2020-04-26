package main

import (
	"fmt"
	"time"
)

func strToTime(s string) time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		panic(err)
	}
	return t
}

func main() {
	var a, b, c, n, h, m int
	var ans time.Time
	fmt.Scan(&a, &b, &c)
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&h, &m)
		t, _ := time.Parse("15:04", fmt.Sprintf("%02d:%02d", h, m))
		t2 := t.Add(time.Duration(b+c) * time.Minute)
		if t2.Before(strToTime("09:00")) {
			ans = t.Add((time.Duration(-a)) * time.Minute)
		}
	}
	fmt.Println(ans.Format("15:04"))
}
