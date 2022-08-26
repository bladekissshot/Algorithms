package Leetcode

func longestCommonSubsequence(text1 string, text2 string) int {

	n1 := len(text1)
	n2 := len(text2)

	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for r := 1; r <= n1; r++ {
		for c := 1; c <= n2; c++ {
			if text1[r-1] == text2[c-1] {
				dp[r][c] = dp[r-1][c-1] + 1
			} else {
				dp[r][c] = max(dp[r][c-1], dp[r-1][c])
			}
		}
	}
	return dp[n1][n2]

}
