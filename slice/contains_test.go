package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainString(t *testing.T) {
	l := []string{"a", "b", "c"}
	assert.True(t, ContainsAny(l, "a"))
	assert.False(t, ContainsAny(l, "d"))
}

func Test_ContainInt(t *testing.T) {
	l := []int{1, 2, 3}
	assert.True(t, ContainsAny(l, 1))
	assert.False(t, ContainsAny(l, 5))
}

func TestHasSameItem(t *testing.T) {
	l1 := []int{1, 2, 3}
	l2 := []int{3, 5}
	l3 := []int{7, 8}
	assert.True(t, HasSameItem(l1, l2))
	assert.False(t, HasSameItem(l1, l3))
}
