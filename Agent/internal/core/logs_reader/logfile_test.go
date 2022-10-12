package logs_reader

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	logfile   = "../../../test/test_log.log"
	freq_test = 3
)

func Test_ReadOldEvents(t *testing.T) {
	lf, err := NewLogFile(logfile)
	require.NoError(t, err)

	events := make(chan string)
	errors := make(chan error)
	done := make(chan struct{})

	defer close(events)
	defer close(errors)
	defer close(done)
	defer lf.File.Close()

	go lf.ReadOldEvents(events, errors, done)

	cnt := 1

	for cnt != 0 {
		select {
		case event := <-events:
			require.NotEmpty(t, event)
		case err := <-errors:
			require.NoError(t, err)
		case <-done:
			cnt--
		}
	}
}

func Test_ReadNewEvents(t *testing.T) {
	lf, err := NewLogFile(logfile)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	events := make(chan string)
	errors := make(chan error)

	defer close(events)
	defer close(errors)

	go lf.ReadNewEvents(ctx, events, errors, freq_test)

loop:
	for {
		select {
		case event := <-events:
			require.NotEmpty(t, event)
		case err := <-errors:
			require.ErrorIs(t, err, context.DeadlineExceeded)
			break loop
		}
	}
}
