func longestPalindrome(s string) string {

    l := len(s)
    if l<=1 {
        return s
    }
    dp := make([][]bool,l)
    res := s[0:1]

    for i:=0;i<l;i++{
        dp[i] = make([]bool,l)
        dp[i][i] = true
        for j:=0;j<i;j++{
            if s[i]==s[j] && (i-j<=2||dp[j+1][i-1]){
                dp[j][i] = true
            }else{
                dp[j][i] = false
            }

            if dp[j][i] && i-j+1 > len(res)  {
                res = s[j:i+1]
            }
        }
    }

    return res

}
