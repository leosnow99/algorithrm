package xiaohao

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	sqrt := mySqrt(4)
	fmt.Println(sqrt)
}

func TestRadius(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 4}
	radius := findRadius(a, b)
	fmt.Println(radius)
}

func TestAbs(t *testing.T) {
	i := abs(-1)
	fmt.Println(i)

}

func TestThreeSum(t *testing.T) {
	tem := []int{-1, 0, 1, 2, -1, -4}
	sum := threeSum(tem)
	fmt.Println(sum)
}
