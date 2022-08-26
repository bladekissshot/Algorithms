package Leetcode

import "sort"

type intslice [][]int

func (s intslice) Len() int      { return len(s) }
func (s intslice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s intslice) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	} else if s[i][0] == s[j][0] {
		return s[i][1] > s[j][1]
	} else {
		return false
	}
}

func maxEnvelopes(envelopes [][]int) int {

	n := len(envelopes)
	if n <= 1 {
		return 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Sort(intslice(envelopes))
	dp := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}
	return res

}
