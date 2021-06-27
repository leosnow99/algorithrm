package dp

/**
假设有排成一行的N 个位置，记为1~N，N 一定大于或等于2。开始时机器人在其中的M
位置上（M 一定是1～N 中的一个），机器人可以往左走或者往右走，如果机器人来到1 位置，
那么下一步只能往右来到2 位置；如果机器人来到N 位置，那么下一步只能往左来到N-1 位置。
规定机器人必须走K 步，最终能来到P 位置（P 也一定是1～N 中的一个）的方法有多少种。给
定四个参数N、M、K、P，返回方法数。
*/

// N : 位置为1 ~ N，固定参数
// cur : 当前在cur 位置，可变参数
// rest : 还剩res 步没有走，可变参数
// P : 最终目标位置是P，固定参数
func walk(n, cur, rest, p int) int {
	// 没有剩余步数了，当前的cur位置就是最后的位置， 判断当前位置和目标位置是否相同
	if rest == 0 {
		if cur == p {
			return 1
		}
		return 0
	}

	// 如果当前位置在第一位，只能从1走向2
	if cur == 1 {
		return walk(n, 2, rest-1, p)
	}
	// 如果当前位置在最后一位，只能向左走一步
	if cur == n {
		return walk(n, n-1, rest-1, p)
	}

	// 此时可以向左走也可以向右走
	return walk(n, cur-1, rest-1, p) + walk(n, cur+1, rest-1, p)

}

//N: 总共有多少个位置
//M: 当前位置
//K: 可以移动次数
//P: 目标位置
func ways1(n, m, k, p int) int {
	if n < 2 || k < 1 || m > n || m < 1 || p < 1 || p > n {
		return 0
	}
	return walk(n, m, k, p)
}

func ways2(n, m, k, p int) int {
	// 参数无效直接返回
	if n < 2 || k < 1 || m > n || m < 1 || p < 1 || p > n {
		return 0
	}

	dp := make([][]int, k+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	dp[0][p] = 1
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			if j == 1 {
				dp[i][j] = dp[i-1][2]
			} else if j == n {
				dp[i][j] = dp[i-1][n-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]
			}
		}
	}
	return dp[k][m]
}

// 在方案二之上，进行空间压缩
func ways3(n, m, k, p int) int {
	// 参数无效直接返回
	if n < 2 || k < 1 || m > n || m < 1 || p < 1 || p > n {
		return 0
	}

	dp := make([]int, n+1)
	dp[p] = 1
	for i := 1; i < k; i++ {
		leftUP := dp[1]
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if j == 1 {
				dp[j] = dp[j+1]
			} else if j == n {
				dp[j] = leftUP
			} else {
				dp[j] = leftUP + dp[j+1]
			}
			leftUP = temp
		}
	}
	return dp[m]
}

/**
解决一个问题，如果没有想到显而易见的求解策略（比如数学公式、贪心策略等，都是显
而易见的求解策略），那么就想如何通过尝试的方式找到答案，一旦写出了好的尝试函数，后面
的优化过程全是固定套路。下面介绍本题如何从暴力递归优化成动态规划的解法。暴力递归优
化成动态规划时，首先根据walk 函数的含义结合题意，分析整个递归过程是不是无后效性的。
代码面试中出现的需要利用尝试解法解决的问题，绝大多数都是无后效性的，有后效性的递归
过程在面试中出现的情况极其罕见，这是一个真实情况，本书不再详述。但是分析一个递归过
程是不是无后效性的，依然非常重要，可以帮我们确定这个暴力递归能不能改成动态规划。所
谓无后效性，是指一个递归状态的返回值与怎么到达这个状态的路径无关。
*/
