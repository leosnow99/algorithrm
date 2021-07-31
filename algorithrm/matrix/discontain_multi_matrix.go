package matrix

/**
不包含本位置值的累乘数组
【题目】
给定一个整型数组arr，返回不包含本位置值的累乘数组。
例如，arr=[2,3,1,4]，返回[12,8,24,6]，即除自己外，其他位置上的累乘。
*/

func product1(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}

	count, all := 0, 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			all *= arr[i]
		} else {
			count++
		}
	}

	res := make([]int, len(arr))
	// 如果数组中不含0，则设置res[i]=all/arr[i（] 0≤i≤n）即可
	if count == 0 {
		for i := range arr {
			res[i] = all / arr[i]
		}
	}

	// 如果数组中有1个0，对唯一的arr[i]==0的位置令res[i]=all，其他位置上的值都是0即可
	if count == 1 {
		for i, number := range arr {
			if number == 0 {
				res[i] = all
			}
		}
	}

	// 如果数组中0 的数量大于1，那么res 所有位置上的值都是0。
	return res
}

func product2(arr []int) []int {
	if len(arr) < 2 {
		return nil
	}

	res := make([]int, len(arr))
	res[0] = arr[0]

	for i, number := range arr {
		res[i] = res[i-1] * number
	}

	tmp := 1
	for i := len(arr); i > 0; i-- {
		res[i] = res[i-1] * tmp
		tmp *= arr[i]
	}
	res[0] = tmp

	return res
}
