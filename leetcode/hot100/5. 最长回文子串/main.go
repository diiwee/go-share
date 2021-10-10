package main

/*****************题目描述***********************
5. 最长回文子串

给你一个字符串 s，找到 s 中最长的回文子串。

*/

/*****************方法一***********************
1.回文串也是面试高频问题，个人觉得这类问题实际意义不大
2.回文字符串的核心就是用双指针（PS:中心扩展算法）
3.解决该类问题的核心思想：从中间开始往俩边扩散来判断
*/

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		ans = maxString(s1, ans)
		ans = maxString(s2, ans)
	}

	return ans

}

//辅助函数判断2个字符串长度
func maxString(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

//辅助函数中心扩展 注意边界问题
func palindrome(s string, left int, right int) string {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]

}
