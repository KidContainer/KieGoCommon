package kiecommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSet(t *testing.T) {
	m := NewSet[int]()
	assert.Equal(t, m.Size(), 0, "size equals to 0")
	m.Add(1)
	assert.Equal(t, m.Contain(1), true, "contain 1")
	assert.Equal(t, m.Contain(2), false, "doesn't contain 2")
	assert.Equal(t, m.Size(), 1, "size equals to 1")
	m.Remove(1)
	assert.Equal(t, m.Contain(1), false, "doesn't contain 1")
	assert.Equal(t, m.Contain(2), false, "doesn't contain 2")
	assert.Equal(t, m.Size(), 0, "size equals to 0")

	m.Reserve(1000)
	assert.Equal(t, m.Size(), 0, "size equals to 0")
}

func TestSyncSet(t *testing.T) {
	m := NewSyncSet[int]()
	m.Add(1)
	assert.Equal(t, m.Contain(1), true, "contain 1")
	assert.Equal(t, m.Contain(2), false, "doesn't contain 2")
	m.Remove(1)
	assert.Equal(t, m.Contain(1), false, "doesn't contain 1")
	assert.Equal(t, m.Contain(2), false, "doesn't contain 2")
}