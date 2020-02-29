package main

import (
	"fmt"
)

type Dice struct {
	Nums []int
}

func (d *Dice) rotate(direction string) {
	switch direction {
	case "W":
		d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[2], d.Nums[5], d.Nums[3], d.Nums[0]
	case "E":
		d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[3], d.Nums[0], d.Nums[2], d.Nums[5]
	case "N":
		d.Nums[0], d.Nums[4], d.Nums[5], d.Nums[1] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
	case "S":
		d.Nums[5], d.Nums[1], d.Nums[0], d.Nums[4] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
	}
}

func main() {
	dice := Dice{
		make([]int, 6),
	}
	for i := 0; i < 6; i++ {
		var n int
		fmt.Scan(&n)
		dice.Nums[i] = n
	}
	var directionOrder string
	fmt.Scan(&directionOrder)
	for _, c := range directionOrder {
		dice.rotate(string(c))
	}
	fmt.Println(dice.Nums[0])
}
