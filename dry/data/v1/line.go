package v1

import "github.com/youngzhu/golab/dry/data"

// v0
// 问题：根据DRY原则，length这个属性没有必要，起点和终点就决定了长度
// 可以通过计算获得
// v1
// 问题：必要时，可以通过缓存数据来避免重复进行代价高昂的运算

type Line struct {
	start data.Point
	end   data.Point
}

func (ln Line) Length() float64 {
	return ln.start.DistanceTo(ln.end)
}
