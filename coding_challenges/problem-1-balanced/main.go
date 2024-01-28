package main

import "fmt"


const MOD = 1000000007

func countBalancedWords(n, d int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, 26)
    }

    // Base case: there is one balanced word of length 1 ending with each character
    for j := 0; j < 26; j++ {
        dp[1][j] = 1
    }
	fmt.Println(dp)
    // Dynamic programming to fill the dp array
    for i := 2; i <= n; i++ {
        for j := 0; j < 26; j++ {
            for k := 1; k <= d; k++ {
				//fmt.Println("iter")
				//fmt.Println(j,k)
                if j+k < 26 {
                    dp[i][j] = (dp[i][j] + dp[i-1][j+k]) % MOD
                }
                if j-k >= 0 {
                    dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % MOD
                }
            }
        }
    }

    // Sum up the counts for all characters to get the total number of balanced words
    result := 0
    for j := 0; j < 26; j++ {
        result = (result + dp[n][j]) % MOD
    }
	fmt.Println(dp)

    return result
}


func main(){
	fmt.Println("Hello world")
	result := countBalancedWords(3,1)
	fmt.Println(result)
}