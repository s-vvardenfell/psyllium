package linux_agent

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	chanCap = 0
	freq    = 3
)

var (
	ErrLogFileListEmpty = errors.New("log file list is empty")
)

type LogsReader struct {
	events   chan string
	errors   chan error
	done     chan int
	logFiles []LogFile
}

// TODO: return interface-type?
func NewLogsReader(logfiles []string) (*LogsReader, error) {
	if len(logfiles) == 0 {
		return nil, ErrLogFileListEmpty
	}

	lgf := make([]LogFile, 0, len(logfiles))

	for i := range logfiles {
		lgf = append(lgf, LogFile{
			FileName: logfiles[i],
		})
	}

	return &LogsReader{
		// events:   make(chan<- core.Event, chanCap),
		events:   make(chan string, chanCap),
		errors:   make(chan error),
		done:     make(chan int),
		logFiles: lgf,
	}, nil
}

func (l *LogsReader) Work() {
	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].readOldEvents(l.events, l.errors, l.done)
		}(i)
	}

	// TODO: check possible impl
	// open chanels here
	// close in readOldEvents
	// range over channels in GetEvents

	l.GetEvents()

	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second) //TODO WithCancel / DeadLine
	// defer cancel()

	for i := range l.logFiles {
		go func(i int) {
			l.logFiles[i].readNewEvents(ctx, l.events, l.errors)
		}(i)
	}

	l.GetEvents()
}

func (l *LogsReader) GetEvents() {
	read := len(l.logFiles)

	for {
		select {
		case event := <-l.events:
			if event != "" {
				fmt.Printf("GOT EVENT: %s", event)
			}
		case err := <-l.errors:
			fmt.Printf("GOT ERROR: %s", err)
			return
		case d := <-l.done:
			if read -= d; read == 0 { // decrement num of files to read
				fmt.Println("All files was read succesfully")
				return
			}
		}
	}
}
