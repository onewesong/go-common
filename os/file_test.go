package os

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExist(t *testing.T) {
	assert.True(t, IsFileExist("./seq.txt"))
	assert.True(t, IsFileExist("/dev/null"))
	assert.False(t, IsFileExist("/dev/seq.txt"))
}
