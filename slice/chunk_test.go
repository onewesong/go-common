package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunks(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	size := 3
	chunks := ChunksSlice(l, size)
	assert.Equal(t, len(chunks), 4)
	t.Logf("%+v", chunks)
}
