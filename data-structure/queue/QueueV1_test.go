package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueV1(t *testing.T) {
	q := CreateQueueV1()
	require.True(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 0)
	require.Equal(t, q.head, 0)
	require.Equal(t, q.tail, 0)
	require.Equal(t, len(q.data), 0)

	q.Enqueue(1)
	require.False(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 1)

	q.Enqueue('s')
	require.False(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 2)

	q.Enqueue("hello")
	require.False(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 3)

	v := q.Dequeue()
	require.NotNil(t, v)
	require.False(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 2)

	q.Clear()
	require.True(t, q.IsEmpty())
	require.Equal(t, q.Lenght(), 0)
	require.Equal(t, q.head, 0)
	require.Equal(t, q.tail, 0)
	require.Equal(t, len(q.data), 0)
}
