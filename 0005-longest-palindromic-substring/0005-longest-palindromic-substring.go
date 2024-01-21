
func longestPalindrome(s string) string {
    n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1
    
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])

			if dp[i][j] && length > maxLength {
				start = i
				maxLength = length
			}
		}
	}

	return s[start : start+maxLength]
}