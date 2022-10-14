package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	// int
	assert.True(t, InSlice([]int{1, 2, 3}, 3))
	assert.False(t, InSlice([]int{1, 2, 5}, 3))

	// uint8
	assert.True(t, InSlice([]uint8{1, 2, 3}, 3))
	assert.False(t, InSlice([]uint8{1, 2, 5}, 3))

	// string
	assert.True(t, InSlice([]string{"a", "b", "c"}, "b"))
	assert.False(t, InSlice([]string{"a", "b", "d"}, "c"))

	// ptr
	i1 := 1
	i2 := 2
	i3 := 3
	i4 := 4
	i5 := 4
	i6 := 6
	assert.True(t, InSlice([]*int{&i1, &i2, &i3, &i4}, &i1))
	assert.False(t, InSlice([]*int{&i1, &i2, &i3, &i4}, &i6))
	assert.False(t, InSlice([]*int{&i1, &i2, &i3, &i4}, &i5))
}
