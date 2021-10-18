package other

import (
	"math"
)

/**
判断一个点是否在三角形内部

在二维坐标系中，所有的值都是double 类型，那么一个三角形可以由3 个点来代表，给定3 个点代表的三角形，
再给定一个点（x,y），判断（x,y）是否在三角形中。
*/

// 解法一:
// 	如果点O 在三角形ABC 内部，那么有面积ABC=面积ABO+面积BCO+面积CAO。
// 	如果点O 在三角形ABC 外部，那么有面积ABC<面积ABO+面积BCO+面积CAO。
func getSideLength(x1, y1, x2, y2 float64) float64 {
	a := math.Abs(x1 - x2)
	b := math.Abs(y1 - y2)

	return math.Sqrt(a*a + b*b)
}

func getArea(x1, y1, x2, y2, x3, y3 float64) float64 {
	side1Len := getSideLength(x1, y1, x2, y2)
	side2Len := getSideLength(x1, y1, x3, y3)
	side3Len := getSideLength(x2, y2, x3, y3)
	p := (side1Len + side2Len + side3Len) / 2

	return math.Sqrt((p - side1Len) * (p - side2Len) * (p - side3Len) * p)
}

func isInside1(x1, y1, x2, y2, x3, y3, x, y float64) bool {
	area1 := getArea(x1, y1, x2, y2, x, y)
	area2 := getArea(x1, y1, x3, y3, x, y)
	area3 := getArea(x3, y3, x2, y2, x, y)
	allArea := getArea(x1, y1, x2, y2, x3, y3)

	return allArea == (area1 + area2 + area3)
}

// 如果点O 在三角形ABC 中，那么从三角形的一点出发，逆时针走过所有边的过程中，点O
// 始终都在走过边的左侧。
func isInside2(x1, y1, x2, y2, x3, y3, x, y float64) bool {
	// 如果三角形的点不是逆时针输入，改变一下顺序
	if crossProduct(x3-x1, y3-y1, x2-x1, y2-y1) >= 0 {
		x2, y2 = x3, y3
	}

	if crossProduct(x2-x1, y2-y1, x-x1, y-y1) < 0 {
		return false
	}

	if crossProduct(x3-x2, y3-y2, x-x2, y-y2) < 0 {
		return false
	}

	if crossProduct(x1-x3, y1-y3, x-x3, y-y3) < 0 {
		return false
	}

	return true
}

func crossProduct(x1, y1, x2, y2 float64) float64 {
	return x1*y2 - x2*y1
}
