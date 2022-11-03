package slice

import "github.com/onewesong/go-common/maps"

// Position 判断列表某个元素的位置(泛型支持)
func Position[T comparable](list []T, item T) int {
	for pos, v := range list {
		if v == item {
			return pos
		}
	}
	return -1
}

func Uniq[T comparable](list []T) []T {
	m := make(map[T]bool, len(list))
	for _, i := range list {
		m[i] = true
	}
	return maps.Keys(m)
}

func ToMap[T comparable](list []T) map[T]bool {
	m := make(map[T]bool, len(list))
	for _, i := range list {
		m[i] = true
	}
	return m
}
