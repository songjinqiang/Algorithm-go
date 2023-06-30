package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 哈希集合用于存储窗口中的字符
	set := make(map[byte]bool)
	n := len(s)

	// 右指针，用于移动窗口的右边界
	right := 0
	// 最长子串的长度
	maxLen := 0

	for left := 0; left < n; left++ {
		// 左指针向右移动，需要移除左指针对应的字符
		if left != 0 {
			delete(set, s[left-1])
		}

		// 移动右指针，不断扩张窗口
		for right < n && !set[s[right]] {
			set[s[right]] = true
			right++
		}

		// 更新最长子串的长度
		maxLen = max(maxLen, right-left)

		// 如果右指针已到达字符串末尾，则无需再继续遍历
		if right == n {
			break
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	length := lengthOfLongestSubstring(s)
	fmt.Println(length) // 输出: 3
}
