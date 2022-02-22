package average

/* 求平均数 多种算法 */

// 变形计算
func averageOne(a, b uint64) uint64 {
	var high, low uint64
	if a > b {
		high = a
		low = b
	} else {
		low = a
		high = b
	}
	differ := high - low
	// or low + differ/2 - differ&1
	return high - differ/2 - differ&1
}

// 除法前置方案
func averageTwo(a, b uint64) uint64 {
	return a/2 + b/2 + a&b&1
}

// 不考虑精度 直接计算
func averageThree(a, b uint64) uint64 {
	return (a + b) / 2
}
