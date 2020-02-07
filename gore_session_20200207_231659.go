package main

import (
	"github.com/k0kubun/pp"
	"time"
)

func __gore_p(xs ...interface{}) {
	for _, x := range xs {
		pp.Println(x)
	}
}
func main() {
	t := time.Now()
	_, _, _ = t.Date()
	_ = t.Year()
	_ = time.Now().Year()
	_, _, _ = time.Now().Date()
	_ = time.Now().Day()
	_ = time.Now().Hour()
	_ = time.Now().Minute()
	_, _, _ = time.Now().Date()
}
