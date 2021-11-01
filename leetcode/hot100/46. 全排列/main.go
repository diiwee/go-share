package main

func permute(nums []int) [][]int {

	ans := [][]int{}

	if len(nums) < 1 {
		return ans
	}

	backtrack := func(nums []int, path []int) {}

	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 { //满足结束条件
			tmp := make([]int, len(path))
			//因为该 path 变量存的是地址引用，结束当前递归时，将它加入 res 后，
			//还要进入别的递归分支继续搜索，还要继续传递该 path 给别的递归调用，它所指向的内存空间还要继续被操作，
			//所以 res 中的 path 的内容会被改变，这就不对。所以要弄一份当前的拷贝，放入 res，
			//这样后续对 path 的操作，就不会影响已经放入 res 的内容。
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		for i := 0; i < len(nums); i++ { //循环递归
			//做选择
			curr := nums[i]
			path = append(path, curr) //记录路径

			nums = append(nums[:i], nums[1+i:]...) //把当前节点移除
			backtrack(nums, path)                  //递归 ，进入下一层决策树

			//撤销选择
			nums = append(nums[:i], append([]int{curr}, nums[i:]...)...) //nums复原
			path = path[:len(path)-1]                                    //path复原

		}
	}

	backtrack(nums, []int{})

	return ans
}
