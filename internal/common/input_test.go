package common

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInputFilename(t *testing.T) {
	_, fn, _, ok := runtime.Caller(0)
	require.True(t, ok)

	n := 1
	expected := filepath.Join(filepath.Dir(fn), fmt.Sprintf("../../input/%d.txt", n))
	require.Equal(t, expected, getInputFilename(n))
}

func TestLoadInputAsync_Success(t *testing.T) {
	require := require.New(t)
	datastream := loadInputAsync(context.Background(), 1, ChannelSizeDefault)
	for data := range datastream {
		require.NoError(data.Err)
		require.NotEmpty(data.Content)
	}
}

func TestLoadInputAsync_InvalidFileLocation(t *testing.T) {
	datastream := loadInputAsync(context.Background(), 0, ChannelSizeDefault)

	// Receive data indicating the error
	data, ok := <-datastream
	require.Nil(t, data.Content)
	require.Error(t, data.Err)
	require.True(t, ok)

	// Channel notifying closure
	data, ok = <-datastream
	require.False(t, ok)
	require.Nil(t, data)
}

func TestLoadInputAsync_Canceled(t *testing.T) {
	require := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	datastream := loadInputAsync(ctx, 1, -1)

	// cancel the context manually, but we do not know when select statement
	// will choose ctx.Done() over out<-
	cancel()
	for data := range datastream {
		require.NoError(data.Err)
		require.NotEmpty(data.Content)
	}
}
