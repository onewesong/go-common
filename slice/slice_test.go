package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	assert.Equal(t, -1, Position([]int{1, 2, 3}, 0))
	assert.Equal(t, 0, Position([]int{1, 2, 3}, 1))
	assert.Equal(t, 1, Position([]int{1, 2, 3}, 2))
}

func TestUniq(t *testing.T) {
	assert.Equal(t, Uniq([]int{1, 1, 1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, Uniq([]string{"1", "1", "1", "2", "3"}), []string{"1", "2", "3"})
}
