package logs_reader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var event = `Oct 17 10:40:31 computer gdm-launch-environment]:` +
	` pam_unix(gdm-launch-environment:session): session opened for user gdm(uid=127) by (uid=0)`

func Test_FormatSysLog(t *testing.T) {
	res, err := FormatSysLog(event)
	require.NoError(t, err)

	require.Equal(t, int64(1666003231), res.DateTime)
	require.Equal(t, "computer", res.Host)
	require.Equal(t, "gdm-launch-environment]", res.Process)
	require.Equal(t,
		"pam_unix(gdm-launch-environment:session) session opened for user gdm(uid=127) by (uid=0)",
		res.Msg)
}
