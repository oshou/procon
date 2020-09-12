package main

func isHappy(n int) bool {
	var (
		sum int
		num int = n
	)
	for sum != 1 {
		for num != 0 {
			tmp := num % 10
			sum += tmp * tmp
			num /= 10
			fmt.Println("sum:",sum)
		}
	}
	return true
}
