package ch26

import "math"

// Hypotenuse 余弦定理
func Hypotenuse(sideA, sideB float64) float64 {
	return math.Sqrt((sideA * sideA) + (sideB * sideB))
}

var (
	cachedHypotenuse         = 0.0
	cachedSideA, cachedSideB = 0.0, 0.0
)

func HypotenuseWithCache(sideA, sideB float64) float64 {
	if hitCache(sideA, sideB) {
		return cachedHypotenuse
	}
	// 没有命中缓存
	// 计算并缓存
	cachedHypotenuse = math.Sqrt((sideA * sideA) + (sideB * sideB))
	cachedSideA = sideA
	cachedSideB = sideB

	return cachedHypotenuse
}

func hitCache(a float64, b float64) bool {
	return (a == cachedSideA && b == cachedSideB) || (a == cachedSideB && b == cachedSideA)
}
