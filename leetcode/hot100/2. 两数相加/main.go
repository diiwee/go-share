package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*****************题目描述***********************
2. 两数相加

给你两个非空 的链表，表示两个非负的整数。
它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字0之外，这两个数都不会以 0开头。
*/

/*****************方法一***********************
1.根据题目描述两数相加，从题目能看出，因为是倒序存储，所以直接吧对应位相加判断进位即可
长度相对较短的链表补0即可
2.这题也是链表操作的基础，希望各位要熟练到信手拈来
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var ans *ListNode
	var head *ListNode
	advance := 0

	for l1 != nil || l2 != nil {
		//L1和L2当前值
		currL1, currL2 := 0, 0
		if l1 != nil {
			currL1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currL2 = l2.Val
			l2 = l2.Next
		}

		sum := currL1 + currL2 + advance

		sum, advance = sum%10, sum/10

		if ans == nil {
			head = &ListNode{Val: sum}
			ans = head
		} else {
			ans.Next = &ListNode{Val: sum}
			ans = ans.Next
		}
	}

	if advance > 0 {
		ans.Next = &ListNode{Val: advance}
	}

	ans = head

	return ans

}

/*********************以下2道题目类似***********************

/*****************题目描述***********************
445. 两数相加 II

给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
*/
/*****************方法一***********************
1.和俩数相加一样，只是顺序颠倒，因此可以先将链表改造成栈
2.剩下的就和俩数相加一样了
*/

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	var s1 []int
	var s2 []int

	for l1 != nil || l2 != nil {

		if l1 != nil {
			s1 = append(s1, l1.Val)
			l1 = l1.Next
		}

		if l2 != nil {
			s2 = append(s2, l2.Val)
			l2 = l2.Next
		}

	}

	lenS1, lenS2 := len(s1), len(s2)

	advancd, sum := 0, 0

	var ans *ListNode

	for i := 0; i < lenS1 || i < lenS2; i++ {
		sum = advancd
		if i < lenS1 {
			sum += s1[lenS1-i-1]
		}

		if i < lenS2 {
			sum += s2[lenS2-i-1]
		}

		sum, advancd = sum%10, sum/10
		node := &ListNode{Val: sum}

		if ans == nil {
			ans = node
		} else {
			node.Next = ans
			ans = node
		}
	}

	if advancd > 0 {
		ans = &ListNode{Val: advancd, Next: ans}
	}

	return ans
}

/*****************题目描述***********************
415. 字符串相加

给定两个字符串形式的非负整数num1 和num2，
计算它们的和并同样以字符串形式返回。
你不能使用任何內建的用于处理大整数的库（比如 BigInteger），
也不能直接将输入的字符串转换为整数形式。
*/

/*****************方法一***********************
1.遍历字符串每个元素对位相加即可，注意Go的类型就行
2.如果对Go的类型不熟悉，多多复习
*/
func addStrings(num1 string, num2 string) string {

	num1Len, num2Len := len(num1), len(num2)

	ans := ""
	sum := 0
	advance := 0

	for i := 0; i < num1Len || i < num2Len; i++ {
		sum = advance

		if i < num1Len {
			sum += int(num1[num1Len-i-1] - '0')
		}

		if i < num2Len {
			sum += int(num2[num2Len-i-1] - '0')
		}

		sum, advance = sum%10, sum/10

		ans = strconv.Itoa(sum) + ans
	}

	if advance > 0 {
		ans = strconv.Itoa(advance) + ans
	}

	return ans

}
