package matrix

import (
	"algorithm/util"
)

/**
容器盛水问题
【题目】
给定一个数组arr，已知其中所有的值都是非负的，将这个数组看作一个容器，请返回容器
能装多少水。
*/

/**
一个简洁的标准。如果现在来到i 位置，只单独考虑i 位置的上方能有几格水。
i 位置上方水的数量 = max{ min{ i 左侧的最大值,i 右侧的最大值 } - arr[i] , 0 }
如果我们依次求出数组中每一个位置上方的水，都累加起来就是答案.
*/
func getWater1(arr []int) (res int) {
	if len(arr) < 3 {
		return 0
	}

	// 0 位置和n-1 位置上方一定没有水，所以不尝试
	for i := 1; i < len(arr)-1; i++ {
		leftMax, rightMax := 0, 0

		// 遍历求i 位置的左侧最大值
		for l := 0; l < i; l++ {
			leftMax = util.Max(leftMax, arr[l])
		}
		// 遍历求i 位置的右侧最大值
		for r := len(arr) - 1; r > i; r-- {
			rightMax = util.Max(rightMax, arr[r])
		}

		res += util.Max(util.Min(leftMax, rightMax)-arr[i], 0)
	}

	return
}

/**
生成和arr 等长的两个数组leftMaxs 和rightMaxs，leftMax[i]的含义是arr[0..i]的最大值，
rightMaxs[i]的含义是arr[i..N-1]的最大值，比如arr=[3,1,5,6,7,6,3]，从左往右遍历生成leftMaxs，
leftMaxs[i]=max{leftMaxs[i-1], arr[i]}，得到leftMaxs=[3,3,5,6,7,7,7]。从右往左遍历生成rightMaxs，
rightMaxs[i]=max{rightMaxs[i+1], arr[i]}，得到rightMaxs=[7,7,7,7,7,6,3]。很明显，遍历两次arr 生
成两个预处理数组的时间复杂度为O(N)，之后对于任何一个i 位置，左侧的最大值就是
leftMax[i-1]，右侧的最大值就是rightMax[i+1]。
*/

func getWater2(arr []int) (res int) {
	if len(arr) < 3 {
		return 0
	}

	leftMaxs := make([]int, len(arr))
	leftMaxs[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		leftMaxs[i] = util.Max(leftMaxs[i-1], arr[i])
	}

	rightMaxs := make([]int, len(arr))
	rightMaxs[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		rightMaxs[i] = util.Max(rightMaxs[i+1], arr[i])
	}

	for i := 1; i < len(arr)-1; i++ {
		res += util.Max(util.Min(leftMaxs[i-1], rightMaxs[i+1]-arr[i]), 0)
	}

	return
}

/**
本题的最优解可以做到时间复杂度为O(N)，额外空间复杂度为O(1)。设置左右两个指针，记为
L 和R，还有两个变量leftMax 和rightMax，初始时L 指向arr[1]的位置，R 指向arr[N-2]的位置，
leftMax=arr[0]，rightMax=arr[N-1]，一共就这4 个变量。求解每一步让L 向右移动或者R 向左移
动，leftMax 表示arr[0..L-1]中的最大值，rightMax 表示arr[R+1..N-1]中的最大值

1）如果leftMax 小于或等于rightMax，此时可以求出L 位置上方的水量。这是因为rightMax是arr[R+1..N-1]的最大值，
	而L 的右侧还有一个未遍历的区域，所以L 右侧最大值一定不会小于rightMax。leftMax 代表L 左侧的最大值，
	此时的假设又是leftMax 小于或等于rightMax，所以可知左侧最大值leftMax 是L 位置的瓶颈。故L 位置上方的水量=Max{leftMax - arr[L],0}。
	然后让L向右移动，在移动之前leftMax 要更新。（leftMax=Max{leftMax, arr[L++]}）
2）如果leftMax 大于rightMax，此时可以求出R 位置上方的水量。解释与情况一同理，R位置上方的水量=max{rightMax - arr[R],0}。
	然后让R 向左移动，在移动之前rightMax 要更新。（rightMax=Max{rightMax, arr[R--]}）
3）每一步都会求出L 或者R 一个位置的水量，把这些水量都累加起来，当L 和R 相遇之后一旦错过（L > R），过程就结束。
*/
func getWater3(arr []int) (res int) {
	if len(arr) < 3 {
		return 0
	}

	leftMax, rightMax := arr[0], arr[len(arr)-1]
	l, r := 1, len(arr)-2

	for l < r {
		if leftMax < rightMax {
			res += util.Max(0, leftMax-arr[l])
			leftMax = util.Max(leftMax, arr[l])
			l++
		} else {
			res += util.Max(0, rightMax-arr[r])
			rightMax = util.Max(rightMax, arr[r])
			r--
		}
	}

	return res
}
