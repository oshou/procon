package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	//var res, tmp int
	var num = 1
	var cnt = 1
	//var ans string
	fmt.Scan(&n)
	//alpha := map[int]string{
	//	1:  "a",
	//	2:  "b",
	//	3:  "c",
	//	4:  "d",
	//	5:  "e",
	//	6:  "f",
	//	7:  "g",
	//	8:  "h",
	//	9:  "i",
	//	10: "j",
	//	11: "k",
	//	12: "l",
	//	13: "m",
	//	14: "n",
	//	15: "o",
	//	16: "p",
	//	17: "q",
	//	18: "r",
	//	19: "s",
	//	20: "t",
	//	21: "u",
	//	22: "v",
	//	23: "w",
	//	24: "x",
	//	25: "y",
	//	26: "z",
	//}

	for {
		fmt.Println(num)
		if n <= num {
			break
		}
		num *= 26
		cnt++
	}
	fmt.Println(cnt)

	for i := cnt; i >= 1; i-- {
		fmt.Println(num / int(math.Pow(26, i)))
	}
	//for {
	//	mod = res / 26
	//	for mod <= 26 {

	//	}
	//	ans += alpha[tmp]
	//	fmt.Println("WIP:", ans)

	//	res = n % 26
	//	if res != 0 {
	//		ans += alpha[res]
	//		fmt.Println(ans)
	//		return
	//	}
	//}
	//fmt.Println(ans)
}
