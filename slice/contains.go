package slice

// ContainsAny 判断列表是否包含某个元素(泛型支持)
func ContainsAny[T comparable](list []T, item T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// HasSameItem 判断两个列表间是否包含相同的元素
func HasSameItem[T comparable](list1, list2 []T) bool {
	for _, v2 := range list2 {
		for _, v1 := range list1 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}
