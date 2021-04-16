package main

import (
	"algorithm/src/dp"
	"fmt"
)

func main() {
	result := dp.MaxCoins2([]int{3, 2, 5})
	fmt.Println(result)

	dp.TestSortedEnvelopes()

	lcse := dp.Lcse("a12sdfk1", "sdfk123j")
	fmt.Println(lcse)
	lcst1 := dp.Lcst2("sdfk213j", "a12sdfk1")
	fmt.Println(lcst1)
}
