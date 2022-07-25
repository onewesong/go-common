package slice

// Index 判断列表某个元素的位置(泛型支持)
func Index[T comparable](list []T, item T) int {
	for pos, v := range list {
		if v == item {
			return pos
		}
	}
	return -1
}
