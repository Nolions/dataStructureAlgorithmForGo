package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := []int{8, 4, 3, 0, 5, 7, 6}
	expected := []int{0, 3, 4, 5, 6, 7, 8}

	insertSort(data)

	assert.Equal(t, expected, data)
}
