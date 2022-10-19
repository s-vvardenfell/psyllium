package logs_reader

import (
	"agent/internal/core"
	"errors"
	"fmt"
	"time"
)

var (
	ErrLogFileListEmpty               = errors.New("log file list is empty")
	resChanCap                        = 100
	evChanCap                         = 0
	deadLine            time.Duration = 60
	freq                              = 3
)

type LogFileReader struct {
	// events   chan core.Event
	// errors   chan error
	// done     chan struct{}
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
		logFiles: lgfs,
	}, nil
}

func (l *LogFileReader) Work(
	events chan<- core.Event, done chan<- struct{}, err chan<- error) {

	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].ReadOldEvents(events, done, err)
		}(i)
	}
}
