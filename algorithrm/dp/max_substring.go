package dp

/**
给定数组 arr，返回arr 的最长递增子序列。
arr=[2,1,5,3,6,4,8,9,7]，返回的最长递增子序列为[1,3,4,8,9]
*/
func list1(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	dp := getDP(arr)
	return generateLTS(arr, dp)
}

func getDP(arr []int) []int {
	dp := make([]int, len(arr))
	for i := range arr {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp
}

func generateLTS(arr, dp []int) []int {
	length := 0
	index := 0
	for i := range dp {
		if length < dp[i] {
			length = dp[i]
			index = i
		}
	}

	lis := make([]int, length)
	length--
	lis[length] = arr[index]
	for i := index; i > 0; i-- {
		if arr[i] < lis[index] && arr[i] == dp[index]-1 {
			length--
			lis[length] = arr[i]
			index = i
		}
	}
	return lis
}

/**
时间复杂度O(NlogN)生成dp 数组的过程是利用二分查找来进行的优化。先生成一个长度为
N 的数组ends，初始时ends[0]=arr[0]，其他位置上的值为0。生成整型变量right，初始时right=0。
在从左到右遍历arr 数组的过程中，求解dp[i]的过程需要使用ends 数组和right 变量，所以这里
解释一下其含义。遍历的过程中，ends[0..right]为有效区，ends[right+1..N-1]为无效区。对有效
区上的位置b，如果有ends[b]==c，则表示遍历到目前为止，在所有长度为b+1 的递增序列中，
最小的结尾数是c。无效区的位置则没有意义。
 */
func getDP2(arr []int) []int {
	dp := make([]int, len(arr))
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	dp[0] = 1
	right, l, r, m := 0, 0, 0, 0

	for i := 1; i < len(arr); i++ {
		l = 0
		for r = right; l < r; {
			m = l + (r-l)>>1
			if ends[m] < arr[i] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if right < l {
			right = l
		}
		ends[l] = arr[i]
		dp[i] = l + 1
	}
	return dp
}

func lis2(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	dp := getDP2(arr)
	return generateLTS(arr, dp)
}
