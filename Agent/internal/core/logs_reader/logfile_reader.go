package logs_reader

import (
	"agent/internal/core"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	ErrLogFileListEmpty               = errors.New("log file list is empty")
	resChanCap                        = 100
	evChanCap                         = 0
	deadLine            time.Duration = 60
	freq                              = 3
)

type LogFileReader struct {
	results  chan core.Event
	events   chan core.Event
	errors   chan error
	done     chan struct{}
	logFiles []LogFile
}

func NewLogsReader(files []string) (*LogFileReader, error) {
	if len(files) == 0 {
		return nil, ErrLogFileListEmpty
	}

	lgfs := make([]LogFile, 0, len(files))

	for i := range files {
		lf, err := NewLogFile(files[i])
		if err != nil {
			return nil, fmt.Errorf("cannot create LogFile with file %s; %w", files[i], err)
		}

		lgfs = append(lgfs, *lf)
	}

	return &LogFileReader{
		results:  make(chan core.Event, resChanCap),
		events:   make(chan core.Event, evChanCap),
		errors:   make(chan error),
		done:     make(chan struct{}),
		logFiles: lgfs,
	}, nil
}

func (l *LogFileReader) Work() <-chan core.Event {
	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].ReadOldEvents(l.events, l.errors, l.done)
		}(i)
	}

	cnt := len(l.logFiles)

loop1:
	for cnt != 0 {
		select {
		case event := <-l.events:
			l.results <- event
		case err := <-l.errors:
			logrus.Error(err)
			break loop1
		case <-l.done:
			cnt--
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), deadLine*time.Second)
	defer cancel()

	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].ReadNewEvents(ctx, l.events, l.errors, freq)
		}(i)
	}

loop2:
	for {
		select {
		case event := <-l.events:
			l.results <- event
		case err := <-l.errors:
			logrus.Error(err)
			break loop2
		}
	}

	return l.results
}
