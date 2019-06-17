package numDecodings

import "strconv"

func numDecodings(s string) int {
	if string(s[0]) == "0" {
		return 0
	}

	dp := []int{1,1}
	count := len(s)
	for i := 1; i < count; i++ {
		if string(s[i-1]) != "0" {
			num := string(s[i-1] + s[i])
			if num >= "1" && num <= "26" {
				if string(s[i]) != "0" {
					dp = append(dp, dp[i-1] + dp[i])
				} else {
					dp = append(dp, dp[i-1])
				}
			} else if string(s[i]) != "0" {
				dp = append(dp, dp[i])
			} else {
				return 0
			}
		} else if string(s[i]) != "0" {
			dp = append(dp, dp[i-1])
		} else {
			return 0
		}
	}

	return dp[count]
}


func numDecoding(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		current := string(s[i-1])
		if current != "0" {
			dp[i] = dp[i] + dp[i-1]
		}

		if i > 1 {
			num, _ := strconv.Atoi(string(s[i-2]))
			intCurrent, _ := strconv.Atoi(current)
			num = num * 10 + intCurrent
			if num == 0 {
				return 0
			}
			if num >= 10 && num <= 26 {
				dp[i] = dp[i] + dp[i-2]
			}
		}
	}

	if len(s) == 0 {
		return 0
	} else {
		return dp[len(s)]
	}
}
