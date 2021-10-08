package main

/*****************题目描述***********************
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

/*****************方法一***********************
1.此类字符串匹配问题，你脑海中应该第一反应滑动窗口算法
2.先来看看滑动窗口的代码框架

	func slidingWindow(s string){
		left,right := 0,0
		windows := map[byte]int{}
		//遍历字符串
		for right<len(s){
			rs := s[right]
			//把右边加入map
			windows[rs]++
			right++

			//缩小窗口
			for windows {
				ls := s[left]
				windows[ls]--
				left++
			}
		}
	}
这就是滑动窗口的核心代码
*/
func lengthOfLongestSubstring(s string) int {

	window := map[byte]int{}

	left, right, ans := 0, 0, 0

	for right < len(s) {
		r := s[right]
		right++
		window[r]++

		//当window[r]>1时 说明 已经遇到重复字符
		for window[r] > 1 {
			l := s[left]
			left++
			window[l]--
		}

		ans = max(ans, right-left)
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
