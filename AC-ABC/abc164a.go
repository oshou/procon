package abc164a

import "fmt"

func IsSafe164a(s, w int) string {
	if s <= w {
		return "unsafe"
	}
	return "safe"
}

func main() {
	var s, w int
	fmt.Scan(&s, &w)
	if s <= w {
		fmt.Println("unsafe")
	} else {
		fmt.Println("safe")
	}
}
