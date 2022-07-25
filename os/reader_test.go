package os

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const fpath = "./seq.txt"

func TestReadLines(t *testing.T) {
	got, err := ReadLines(fpath, 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, got)
}

func TestReadFirstLine(t *testing.T) {
	got, err := ReadFirstLine(fpath)
	assert.NoError(t, err)
	assert.Equal(t, "1", got)
}

func TestReadFirstLineAsInt(t *testing.T) {
	got, err := ReadFirstLineAsInt(fpath)
	assert.NoError(t, err)
	assert.Equal(t, 1, got)
}
