package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueV1(t *testing.T) {
	lrc := Constructor(10)
	lrc.Put(10, 13)
	lrc.Put(3, 17)
	lrc.Put(6, 11)
	lrc.Put(10, 5)
	lrc.Put(9, 10)
	require.Equal(t, lrc.Get(13), -1)
	require.Equal(t, lrc.Get(5), -1)
	lrc.Put(2, 19)
	require.Equal(t, lrc.Get(2), 19)
	require.Equal(t, lrc.Get(3), 17)
	lrc.Put(5, 25)
	require.Equal(t, lrc.Get(8), -1)
	lrc.Put(9, 22)
	lrc.Put(5, 5)
	lrc.Put(1, 30)
	require.Equal(t, lrc.Get(11), -1)
}
