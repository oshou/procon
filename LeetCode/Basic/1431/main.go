package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, candy := range candies {
		if candy >= max {
			max = candy
		}
	}
	ans := make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= max {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
