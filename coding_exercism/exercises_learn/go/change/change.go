package change

import "errors"

func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("negative target value is not allowed")
	}

	dp := make([][]int, target+1)
	for i := range dp {
		dp[i] = make([]int, len(coins))
	}

	for i := 1; i <= target; i++ {
		for j, coin := range coins {
			if i < coin {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-coin][j]
			}

			if i == coin {
				dp[i][j]++
			}
		}
	}

	if dp[target][len(coins)-1] == 0 {
		return nil, errors.New("no combination can add up to target")
	}

	result := make([]int, 0)
	i, j := target, len(coins)-1
	for i > 0 && j >= 0 {
		if i-coins[j] >= 0 && dp[i][j] > dp[i-coins[j]][j] {
			result = append(result, coins[j])
			i -= coins[j]
		} else {
			j--
		}
	}

	return result, nil	
}

