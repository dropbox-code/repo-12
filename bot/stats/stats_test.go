package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStats(t *testing.T) {
	value, err := Get("test")
	require.Error(t, err)
	assert.Equal(t, value, uint(0))

	Increase("test", 2)
	value, err = Get("test")
	require.NoError(t, err)
	assert.Equal(t, value, uint(2))

	IncreaseOne("test")
	value, err = Get("test")
	require.NoError(t, err)
	assert.Equal(t, value, uint(3))

	Set("test", 42)
	value, err = Get("test")
	require.NoError(t, err)
	assert.Equal(t, value, uint(42))

	Increase("test", int64(1))
	Increase("test", int8(1))
	Increase("test", 1)
	value, err = Get("test")
	require.NoError(t, err)
	assert.Equal(t, value, uint(45))

	IncreaseOne("test")
	value, err = Get("test")
	require.NoError(t, err)
	assert.Equal(t, value, uint(46))
}
