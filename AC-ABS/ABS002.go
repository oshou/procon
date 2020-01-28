package main

import (
	"fmt"
	"os"
)

func main(s int) {
	os.Setenv("key", "value")
	fmt.Println(os.Unsetenv("abc"))
	fmt.Println(os.Getpid())
	os.Args("aaa")
}
