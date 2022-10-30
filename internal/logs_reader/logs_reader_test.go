package logs_reader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ReadLog(t *testing.T) {
	lr := NewLogsReader("../../../test/test_log.log")
	go lr.ReadLog(FormatSysLog, 1640984400)

loop:
	for {
		select {
		case ev := <-lr.Events:
			require.NotEmpty(t, ev.DateTime)
			require.NotEmpty(t, ev.Host)
			require.NotEmpty(t, ev.Process)
			require.NotEmpty(t, ev.Msg)
		case err := <-lr.Errors:
			require.NoError(t, err)
		case <-lr.Done:
			break loop
		}
	}
}
