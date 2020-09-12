package main

func romanToInt(s string) int {
	doubles := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	singles := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var i, sum int
	for i <= len(s)-1 {
		if len(s)-i > 1 && doubles[string(s[i:i+2])] != 0 {
			sum += doubles[string(s[i:i+2])]
			i += 2
		} else {
			sum += singles[string(s[i])]
			i += 1
		}
	}
	return sum
}
