package string_type

/**
旋变字符串问题

一个字符串可以分解成多种二叉树结构。如果str 长度为1，认为不可分解；如果str 长度
为N（N>1），左部分长度可以为1~N-1，剩下的为右部分的长度。左部分和右部分都可以按照
同样的逻辑，继续分解。形成的所有结构都是str 的二叉树结构。
*/

func sameTypeSameNumber(str1, str2 []byte) bool {
	if len(str1) != len(str2) {
		return false
	}
	maps := make([]int, 256)
	for _, s := range str1 {
		maps[s]++
	}
	for _, s := range str2 {
		if maps[s] == 0 {
			return false
		}
		maps[s]--
	}
	return true
}

// 返回str1[从L1 开始往右长度为size 的子串]和str2[从L2 开始往右长度为size 的子串] 是否互为旋变字符串
// 在str1 中的这一段和str2 中的这一段一定是等长的，所以只用一个参数size
func processRotatingStr(str1, str2 []byte, L1, L2, size int) bool {
	if size == 1 {
		return str1[L1] == str2[L2]
	}
	// 枚举每一种情况，有一个计算出互为旋变就返回true，都算不出来则返回false
	for leftPart := 1; leftPart < size; leftPart++ {
		if judgeAllow(str1, str2, L1, L2, leftPart, size) {
			return true
		}
	}
	return false
}

// 如果str1[0..i]和str2[0..i]互为旋变，并且str1[i+1..N-1]和str2[i+1..N-1]互为旋变，则str1
// 和str2 互为旋变字符串；如果str1[0..i]和str2[N-i-1..N-1]互为旋变，并且str1[i+1..N-1]和
// str2[0..N-i-2]互为旋变，则str1 和str2 互为旋变字符串。
func judgeAllow(str1, str2 []byte, L1, L2, leftPart, size int) bool {
	b1 := processRotatingStr(str1, str2, L1, L2, leftPart) && processRotatingStr(str1, str2, L1+leftPart,
		L2+leftPart, size-leftPart)
	if b1 {
		return true
	}
	b2 := processRotatingStr(str1, str2, L1, size-leftPart, leftPart) && processRotatingStr(
		str1, str2, L1+leftPart, L2, size-leftPart)
	return b2
}

func isScramble1(str1, str2 string) bool {
	if len(str1) == 0 && len(str2) == 0 {
		return true
	}
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}
	if str1 == str2 {
		return true
	}
	chas1 := []byte(str1)
	chas2 := []byte(str2)
	if !sameTypeSameNumber(chas1, chas2) {
		return false
	}
	return processRotatingStr(chas1, chas2, 0, 0, len(chas1))
}

func isScramble2(str1, str2 string) bool {
	if len(str1) == 0 && len(str2) == 0 {
		return true
	}
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}
	if str1 == str2 {
		return true
	}
	chas1 := []byte(str1)
	chas2 := []byte(str2)
	if !sameTypeSameNumber(chas1, chas2) {
		return false
	}

	N := len(chas1)
	dp := make([][][]bool, N)
	for i := 0; i < N; i++ {
		demo := make([][]bool, N)
		for j := 0; j < N; j++ {
			demo[j] = make([]bool, N+1)
		}
		dp[i] = demo
	}
	for L1 := 0; L1 < N; L1++ {
		for L2 := 0; L2 < N; L2++ {
			if chas1[L1] == chas2[L2] {
				dp[L1][L2][1] = true
			}
		}
	}
	// 第一层for 循环含义是：依次填size=2 层、size=3 层……size=N 层，每一层都是一个二维平面
	// 第二、三层for 循环含义是：在具体的一层中，整个面都要填写，所以用两个for 循环去填一个二维平面
	// L1 的取值范围是[0,N-size]，
	// 因为从L1 出发往右长度为size 的子串，L1 是不能从N-size+1 出发的，这样往右就不够size 个字符
	// L2 的取值范围同理
	// 第4 层for 循环完全是递归函数怎么写，这里就怎么
	for size := 2; size <= N; size++ {
		for L1 := 0; L1 <= N-size; L1++ {
			for L2 := 0; L2 <= N-size; L2++ {
				for leftPart := 1; leftPart < size; leftPart++ {
					if (dp[L1][L2][leftPart] && dp[L1+leftPart][L2+leftPart][size-leftPart]) || (
						dp[L1][L2+size-leftPart][leftPart] && dp[L1+leftPart][L2][size-leftPart]) {
						dp[L1][L2][size] = true
						break
					}

				}
			}
		}
	}
	return dp[0][0][N]
}
