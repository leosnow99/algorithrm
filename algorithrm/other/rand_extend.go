package other

import (
	"math/rand"
)

/**
从5 随机到7 随机及其扩展
【题目】
给定一个等概率随机产生1~5 的随机函数rand1To5 如下：
	public int rand1To5() {
		return (int) (Math.random() * 5) + 1;
	}
除此之外，不能使用任何额外的随机机制，请用rand1To5 实现等概率随机产生1~7 的随机函数rand1To7。

补充问题：给定一个以p 概率产生0，以1-p 概率产生1 的随机函数rand01p 如下：
	public int rand01p() {
		// 可随意改变p
		double p = 0.83;
		return Math.random() < p ? 0 : 1;
	}
除此之外，不能使用任何额外的随机机制，请用rand01p 实现等概率随机产生1~6 的随机函数rand1To6。

进阶问题：给定一个等概率随机产生1~m 的随机函数rand1ToM 如下：
	public int rand1ToM(int m) {
		return (int) (Math.random() * m) + 1;
	}

除此之外，不能使用任何额外的随机机制。有两个输入参数，分别为m和n，请用rand1ToM(m)实现等概率随机产生1~n 的随机函数rand1ToN。
*/
func rand1To5() int {
	return rand.Intn(5) + 1
}

func rand01p() int {
	p := 0.83
	if rand.Intn(100) < int(p*100) {
		return 0
	}

	return 1
}

func rand1ToM(m int) int {
	return rand.Intn(m) + 1
}

/*
1．rand1To5()等概率随机产生1,2,3,4,5。
2．rand1To5()-1 等概率随机产生0,1,2,3,4。
3．(rand1To5()-1)*5 等概率随机产生0,5,10,15,20。
4．(rand1To5()-1)*5+(rand1To5()-1)等概率随机产生0,1,2,3,…,23,24。注意，这
	两个rand1To5()是指独立的两次调用，请不要简化。这是“插空儿”的过程。
5．如果步骤4 产生的结果大于20，则重复进行步骤4，直到产生的结果在0~20 之间。
	同时可以轻易知道出现21~24 的概率会平均分配到0~20 上。这是“筛”过程。
6．步骤5 会等概率随机产生0~20，所以步骤5 的结果再进行%7 操作，就会等概率地随机产生0~6。
7．步骤6 的结果再加1，就会等概率地随机产生1~7。
*/
func rand1To7() int {
	num := 0
	for {
		num = (rand1To5()-1)*5 + rand1To5() - 1
		if num <= 20 {
			break
		}
	}

	return num%7 + 1
}

/*
虽然rand01p 方法以p 的概率产生0，以1-p 的概率产生1，但是rand01p产生01 和10 的概率却都是p(1-p)，
	可以利用这一点来实现等概率随机产生0 和1 的函数。

有了等概率随机产生0 和1 的函数后，再按照如下步骤生成等概率随机产生1~6 的函数：
	1．rand01()方法可以等概率随机产生0 和1。
	2．rand01()*2 等概率随机产生0 和2。
	3．rand01()*2+rand01()等概率随机产生0,1,2,3。注意，这两个rand01()是指独立的两次调用，请不要化简。这是“插空儿”过程。

	步骤3 已经实现了等概率随机产生0~3 的函数，具体请参看如下代码中的rand0To3 方法：
	public int rand0To3() {
		return rand01() * 2 + rand01();
	}
	4．rand0To3()*4+rand0To3()等概率随机产生0,1,2,…,14,15。注意，这两个rand0To3()是指独立的两次调用，请不要简化。这还是“插空儿”过程。
	5．如果步骤4 产生的结果大于11，则重复进行步骤4，直到产生的结果在0~11 之间。
		那么可以知道出现12~15 的概率会平均分配到0~11 上。这是“筛”过程。
	6．因为步骤5 的结果是等概率随机产生0~11，所以用第5 步的结果再进行%6 操作，就会等概率随机产生0~5。
	7．第6 步的结果再加1，就会等概率随机产生1~6。
*/
func rand01() int {
	num := 0
	for {
		num = rand01p()
		if num != rand01p() {
			break
		}
	}

	return num
}

func rand0To3() int {
	return rand01()*2 + rand01()
}

func rand1To6() int {
	num := 0
	for {
		num = rand0To3()*4 + rand0To3()
		if num <= 11 {
			break
		}
	}

	return num%6 + 1
}

func rand1ToN(n, m int) int {
	nMSys := getMSysNum(n-1, m)
	randNum := getRandMSysNumLessN(nMSys, m)
	return getNumFromMSysNum(randNum, m) + 1
}

func getMSysNum(value, m int) []int {
	res := make([]int, 32)
	index := len(res) - 1
	for value != 0 {
		res[index] = value % m
		value = value / m
		index--
	}

	return res
}

func getRandMSysNumLessN(nMsy []int, m int) []int {
	res := make([]int, len(nMsy))
	start := 0
	for nMsy[start] != 0 {
		start++
	}

	index := start
	lastEqual := true
	for index != len(nMsy) {
		res[index] = rand1ToM(m) - 1
		if lastEqual {
			if res[index] > nMsy[index] {
				index = start
				lastEqual = true
				continue
			} else {
				lastEqual = res[index] == nMsy[index]
			}
		}
		index++
	}

	return res
}

func getNumFromMSysNum(mSysNum []int, m int) (res int) {
	for i := 0; i < len(mSysNum); i++ {
		res = res*m + mSysNum[i]
	}

	return
}
