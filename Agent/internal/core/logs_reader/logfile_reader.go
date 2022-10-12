package logs_reader

import (
	"context"
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
	results  chan string
	events   chan string
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
		results:  make(chan string, resChanCap),
		events:   make(chan string, evChanCap),
		errors:   make(chan error),
		done:     make(chan struct{}),
		logFiles: lgfs,
	}, nil
}

func (l *LogFileReader) Work() <-chan string {
	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].ReadOldEvents(l.events, l.errors, l.done)
		}(i)
	}

	cnt := len(l.logFiles)

	for cnt != 0 {
		select {
		case event := <-l.events:
			fmt.Println(event)
			l.results <- event
		case err := <-l.errors:
			fmt.Println(err) //todo
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

loop:
	for {
		select {
		case event := <-l.events:
			fmt.Println(event)
			l.results <- event
		case err := <-l.errors:
			fmt.Println(err) //todo
			break loop
		}
	}

	return l.results
}
