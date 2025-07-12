package main

import "fmt"

// 生成给定数的前缀的偶数或奇数长度的十进制回文数
// 如 123 -> 12321，123321
func makePalindrome(n int64, odd bool) int64 {
	res := n
	if odd {
		n = n / 10
	}
	for n > 0 {
		res = res*10 + (n % 10)
		n /= 10
	}

	return res
}

// 检查一个数字在 base 进制下是否为回文
// 反转数字，然后判断是否相等
func isPalindrome(n, base int64) bool {
	reversed := int64(0)
	temp := n
	for temp > 0 {
		reversed = reversed*base + (temp % base)
		temp /= base
	}
	return reversed == n
}

// 提取所有数字，然后判断是否回文
func isPalindrome2(n int64, base int64) bool {
	var digits []int
	for n > 0 {
		digits = append(digits, int(n%base))
		n /= base
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

// 计算所有小于 n 的十进制和 base-k 进制下都是回文的数之和
func sumPalindrome(n, k int64) int64 {
	sum := int64(0)

	// 生成奇数长度回文
	for i := int64(1); ; i++ {
		p := makePalindrome(i, true)
		if p > n {
			break
		}

		if isPalindrome(p, k) {
			sum += p
		}
	}

	// 生成偶数长度回文
	for i := int64(1); ; i++ {
		p := makePalindrome(i, false)
		if p > n {
			break
		}
		if isPalindrome(p, k) {
			sum += p
		}
	}

	return sum
}

func main() {
	var A, N int64
	fmt.Scan(&A, &N)

	fmt.Println(sumPalindrome(N, A))
}
