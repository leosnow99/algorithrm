package string_type

// 完美洗牌问题
// 给定一个长度为偶数的数组arr，长度记为2×N。前N 个为左部分，后N 个为右部分。arr
//就可以表示为{L1,L2,..,Ln,R1,R2,..,Rn}，请将数组调整成{R1,L1,R2,L2,..,Rn,Ln}的样子。

// 数组的长度为len，调整前的位置是i，返回调整之后的位置
// 下标不从 0 开始，从 1 开始
func modifyIndex(i, lens int) int {
	if i <= (lens / 2) {
		return i * 2
	}
	return 2*(i-lens/2) - 1
}

func modifyIndex2(i, lens int) int {
	return (2 * i) % (lens + 1)
}

// 数组必须不为空，且长度为偶数
func shuffle(arr []int) {
	if len(arr) != 0 || len(arr)&1 == 0 {
		shuffleImpl(arr, 0, len(arr)-1)
	}
}

// 在arr[L，R] 上做完美洗牌的调整
func shuffleImpl(arr []int, L, R int) {
	for R-L+1 > 0 {
		lens := R - L + 1
		base := 3
		k := 1
		// 计算小于或等于lens且距离lens最近的，满足（3^k) <= lens + 1 的数字
		// 也就是找到最大的k，满足 3^k <= lens + 1
		for base <= (lens+1)/3 {
			base *= 3
			k++
		}
		// 当前要解决长度为 base - 1 的块，一半就是再除以2
		half := (base - 1) / 2
		// [L, R] 的中点位置
		mid := L + (R-L)>>1
		// 要旋转的做部分为 [L+half...mid], 右部分为 [mid+1...mid+half]
		// 注意，arr 的下标从 0 开始
		rotated(arr, L+half, mid, mid+half)
		// 旋转完成后，从L开始算，长度为 base - 1 的部分进性下标连续推
		cycles(arr, L, base-1, k)
		// 解决了前 base - 1 的部分，剩下的部分继续处理
		L = L + base - 1
	}
}

// 从start位置开始，向右len的长度这一段做下标连续推
func cycles(arr []int, start, lens, k int) {
	// 找到一个出发位置 trigger， 一共 k 个
	// 每一个 trigger 都进行下标连续推
	// 出发位置是从1开始算的，而数组下标是从 0 开始算的
	for i, trigger := 1, 0; i < k; i++ {
		preValue := arr[trigger+start-1]
		cur := modifyIndex(trigger, lens)
		for cur != trigger {
			tmp := arr[cur+start-1]
			arr[cur+start-1] = preValue
			preValue = tmp
			cur = modifyIndex(cur, lens)
		}
		arr[cur+start-1] = preValue
		trigger *= 3
	}
}

// [L...M]为左部分， [M...R] 为右部分，左右两部分互换
func rotated(arr []int, L, M, R int) {
	reverses(arr, L, M)
	reverses(arr, M+1, R)
	reverses(arr, L, R)
}

func reverses(arr []int, L, R int) {
	for L < R {
		arr[L], arr[R] = arr[R], arr[L]
		L++
		R--
	}
}

// 给定一个数组arr，请将数组调整为依次相邻的数字总是先<=、再>=的关系，并
// 交替下去。比如数组中有五个数字，调整成{a,b,c,d,e}，使之满足a<=b>=c<=d>=e
func wiggleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 假设这个排序的额外空间复杂度是 O(1)
	// sort.Sort(arr)
	if len(arr)&1 == 1 {
		shuffleImpl(arr, 1, len(arr)-1)
	} else {
		shuffleImpl(arr, 0, len(arr)-1)
		for i := 0; i < len(arr); i += 2 {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}
