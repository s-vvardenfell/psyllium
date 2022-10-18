package host_info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetHostInfo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := GetHostInfo(ctx)

	require.NoError(t, err)
	require.NotZero(t, res.Home)
	require.NotZero(t, res.Host)
	require.NotZero(t, res.OS)
	require.NotZero(t, res.Shell)
	require.NotZero(t, res.Username)
}
