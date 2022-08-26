package Leetcode

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(height)
	leftmax := make([]int, n)
	rightmax := make([]int, n)

	leftmax[0] = height[0]
	rightmax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftmax[i] = max(height[i], leftmax[i-1])
		rightmax[n-i-1] = max(height[n-i-1], rightmax[n-i])
	}
	res := 0
	for i := 1; i < n; i++ {
		res += min(leftmax[i], rightmax[i]) - height[i]
	}
	return res

}
