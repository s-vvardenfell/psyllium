package logs_reader

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type LogFile struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func NewLogFile(filename string) (*LogFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot create NewLogFile, %w", err)
	}

	return &LogFile{
		File:    f,
		Scanner: bufio.NewScanner(f),
	}, nil
}

func (l *LogFile) readOldEvents(ctx context.Context, events chan<- string, errs chan<- error) {

}

func (l *LogFile) readNewEvents(ctx context.Context, events chan<- string, errs chan<- error) {
	defer l.File.Close()

}
