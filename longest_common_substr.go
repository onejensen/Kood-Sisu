package sprint

func LongestCommonSubstr(str1, str2 string) string {

	len1, len2 := len(str1), len(str2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	maxLength := 0
	endIndex := 0

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i - 1
				}
			}
		}
	}

	if maxLength > 0 {
		return str1[endIndex-maxLength+1 : endIndex+1]
	}

	return ""
}
