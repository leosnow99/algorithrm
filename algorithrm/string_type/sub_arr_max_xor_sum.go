package string_type

import (
	"algorithm/util"
	"math"
)

/*
子数组的最大异或和
数组异或和的定义：把数组中所有的数异或起来得到的值。
给定一个整型数组arr，其中可能有正、有负、有零，求其中子数组的最大异或和。
*/

func maxXorSubArray1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	eor := make([]int, len(arr))
	eor[0] = arr[0]
	// 生成eor 数组，eor[i]代表arr[0..i]的异或和
	for i := 1; i < len(arr); i++ {
		eor[i] = eor[i-1] ^ arr[i]
	}
	max := math.MinInt32
	// 以j 位置结尾的情况下，每一个子数组最大的异或和
	for j := 0; j < len(arr); j++ {
		// 依次尝试arr[0..j],arr[1..j],...,arr[i..j],...,arr[j..j]
		for i := 0; i <= j; i++ {
			max = util.Max(max, eor[i-1]^eor[j])
		}
	}
	return max
}

func maxXORSubArray2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max, eor := math.MinInt32, 0
	numTire := NewNumTire()
	numTire.Add(0)
	for j := 0; j < len(arr); j++ {
		eor ^= arr[j]
		max = util.Max(max, numTire.MaxXor(eor))
		numTire.Add(arr[j])
	}
	return max
}

// 前缀树的节点类型，每个节点向下只可能有走向0 或1 的路
type node struct {
	nexts []*node
}

func newNode() *node {
	return &node{
		nexts: make([]*node, 2),
	}
}

// NumTire 基于本题，定制前缀树的实现
type NumTire struct {
	head *node
}

func NewNumTire() NumTire {
	return NumTire{
		head: newNode(),
	}
}

// Add 把某个数字newNum 加入到这棵前缀树里
// num 是一个32 位的整数，所以加入的过程一共走32 步
func (t NumTire) Add(newNumber int) {
	cur := t.head
	for move := 31; move >= 0; move-- {
		path := (newNumber >> move) & 1
		if cur.nexts[path] == nil {
			cur.nexts[path] = newNode()
		}
		cur = cur.nexts[path]
	}
}

// MaxXor 给定一个eorJ，eorJ 表示eor[j]，即以j 位置结尾的情况下，arr[0..j]的异或和
// 因为之前把eor[0]，eor[1]，…，eor[j-1]都加入了前缀树，所以可以选择出一条最优路径
// maxXor 方法就是把最优路径找到，并且返回eor[j]与最优路径结合之后得到的最大异或和
func (t NumTire) MaxXor(eorJ int) (res int) {
	cur := t.head
	for move := 31; move >= 0; move-- {
		path := (eorJ >> move) & 1
		best := path ^ 1
		if move == 31 {
			best = path
		}
		if cur.nexts[path] == nil {
			best = path ^ 1
		}
		res |= (path ^ best) << move
		cur = cur.nexts[path]
	}
	return
}
