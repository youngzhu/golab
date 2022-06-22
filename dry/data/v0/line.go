package v0

import "github.com/youngzhu/golab/dry/data"

// v0
// 问题：根据DRY原则，length这个属性没有必要，起点和终点就决定了长度
// 可以通过计算获得

type Line struct {
	start  data.Point
	end    data.Point
	length float64
}
