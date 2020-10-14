package a

import "fmt"

func FrogA(n int, h []int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Println("dp[", i, "]:", dp[i])
			continue
		}
		if i == 1 {
			dp[i] = dp[i-1] + (h[i] - h[i-1])
			fmt.Println("dp[", i, "]:", dp[i])
			continue
		}
		dp[i] = min(dp[i-1]+(h[i]-h[i-1]), dp[i-2]+(h[i]-h[i-2]))
		fmt.Println("dp[", i, "]:", dp[i])
	}
	return dp[len(h)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
