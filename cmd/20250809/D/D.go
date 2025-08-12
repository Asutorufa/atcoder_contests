package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	var S string

	fmt.Fscan(br, &N, &S)

	dp := make([][2]int, N+1)

	// d[i][0] 表示当前子串的0是偶数的字串的个数
	// d[i][1] 表示当前子串的0是奇数的字串的个数

	for i := 1; i <= N; i++ {
		if S[i-1] == '0' {
			// 当当前字符是 '0'，长度为 1 的子串 "0" 中 0 的个数是 1（奇数），所以它属于 dp[i][1]（奇数个 0 的子串数量）要加 1。
			dp[i][1] = 1
		} else {
			// 当当前字符是 '1'，长度为 1 的子串 "1" 中 0 的个数是 0（偶数），所以它属于 dp[i][0]（偶数个 0 的子串数量）要加 1。 (0个也是偶数，同时'1'也是最终结果美しい文字列)
			dp[i][0] = 1
		}

		if S[i-1] == '0' {
			// 0时，奇偶翻转
			dp[i][0] += dp[i-1][1]
			dp[i][1] += dp[i-1][0]
		} else {
			// 1时，奇偶不变
			dp[i][0] += dp[i-1][0]
			dp[i][1] += dp[i-1][1]
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += dp[i][0]
	}

	fmt.Println(ans)
}
