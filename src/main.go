package main

import "fmt"

func main() {
	arr := make([]int, 3)
	fmt.Println("len: ", len(arr))
	fmt.Println("cap: ", cap(arr))
	var demo [3][3]int
	demo[1][2] = 1
}
