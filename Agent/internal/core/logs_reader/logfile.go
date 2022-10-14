package logs_reader

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type Event struct {
	FileName string `json:"filename"`
	Event    string `json:"event"`
	Ts       int64  `json:"ts"`
}

type LogFile struct {
	File     *os.File
	Reader   *bufio.Reader
	fileName string //for dev p-s, will be removed
}

func NewLogFile(filename string) (*LogFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot create NewLogFile, %w", err)
	}

	return &LogFile{
		File:     f,
		Reader:   bufio.NewReader(f),
		fileName: filename,
	}, nil
}

// ReadOldEvents reads file contents line by line and send result and error to channels;
// stops when file is completely read
func (l *LogFile) ReadOldEvents(events chan<- Event, errs chan<- error, done chan<- struct{}) {
	for {
		line, err := l.Reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			errs <- fmt.Errorf("error while reading file in ReadOldEvents, %w", err)
			break
		}

		events <- Event{l.fileName, line, time.Now().Unix()}
	}

	done <- struct{}{}
}

// ReadNewEvents read new lines in file and send result and error to channels;
// stops if context is done or error occurs
func (l *LogFile) ReadNewEvents(
	ctx context.Context, events chan<- Event, errs chan<- error, freq int) {
	defer l.File.Close()

	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ctx.Done():
			errs <- ctx.Err()
		case <-ticker.C:
			line, err := l.Reader.ReadString('\n')

			if err != nil {
				if err == io.EOF {
					continue
				}

				errs <- fmt.Errorf("error while reading file in ReadNewEvents, %w", err)
				return
			}

			events <- Event{l.fileName, line, time.Now().Unix()}
		}
	}
}
