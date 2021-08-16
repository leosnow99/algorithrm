package other

import (
	"math"
)

/**
判断一个点是否在矩形内部
【题目】
在二维坐标系中， 类型，:那么一个矩形可以由4 个点来代表，（x1,y1）
为最左的点、（x2,y2）为最上的点、（x3,y3）为最下的点、（x4,y4）为最右的点。给定4 个点代
表的矩形，再给定一个点（x,y），判断（x,y）是否在矩形中。
*/

func inside(x1, y1, x4, y4, x, y float64) bool {
	if x <= x1 {
		return false
	}

	if x >= x4 {
		return false
	}

	if y >= y1 {
		return false
	}

	if y <= y4 {
		return false
	}

	return true
}

/*
因为矩形的边不是平行于x 轴就是平行于y 轴，所以判断该点是否完全在矩形的左侧、右侧、上侧或下侧，如果都不是，就一定在其中。如果矩形的边不平行
于坐标轴呢？也非常简单，就是高中数学的知识，通过坐标变换把矩阵转成平行的情况，在旋转时所有的点跟着转动就可以。旋转完成后，再用上面的方式进行判断
*/
func isInside(x1, y1, x2, y2, x3, y3, x4, y4, x, y float64) bool {
	if y1 == y2 {
		return inside(x1, y1, x4, y4, x, y)
	}

	l := math.Abs(y4 - y3)
	k := math.Abs(x4 - x3)
	s := math.Sqrt(k*k + l*l)
	sin := l / s
	cos := k / s
	x1R := cos*x1 + sin*y1
	y1R := -x1*sin + y1*cos
	x4R := cos*x4 + sin*y4
	y4R := -x4*sin + y4*cos
	xR := cos*x + sin*y
	yR := -x*sin + y*cos

	return inside(x1R, y1R, x4R, y4R, xR, yR)
}
