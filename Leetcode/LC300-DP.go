func lengthOfLIS(nums []int) int {

    if len(nums)<1 {
        return 0
    }
    dp := make([]int,len(nums))

    max := func(a,b int) int{
        if a>b{
            return a
        }else{
            return b
        }
    }

    for i:=0;i<len(nums);i++{
        dp[i]=1
        for j:=0;j<i;j++{
            if nums[i]>nums[j]{
                dp[i] = max(dp[i],dp[j]+1)
            }
        }
    }
    res := math.MinInt32
    for _, k := range dp{
        if k>res {
            res = k
        }
    }

    return res

}
