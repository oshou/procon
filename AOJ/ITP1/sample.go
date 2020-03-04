package main

import (
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open(os.Args[1])
	w, _ := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	buf := make([]byte, 5)
	for {
		n, _ := r.Read(buf)
		fmt.Println(string(buf[:n]), n)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
}
