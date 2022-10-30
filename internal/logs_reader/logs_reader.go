package logs_reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/s-vvardenfell/psyllium/internal/core"
)

var (
	authLog = regexp.MustCompile(
		`([\w]{3,4}\s[\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2}) ([\w|\d|-]{1,40}) (.+)`)
)

type LogsReader struct {
	filename string
	Events   chan *core.Event
	Errors   chan error
	Done     chan struct{}
}

// TODO передать в конструктор LogReader'у каналы для отправки
// ведь LogReader отвечает за 1 файл
func NewLogsReader(filename string) *LogsReader {
	return &LogsReader{
		filename: filename,
		Events:   make(chan *core.Event),
		Errors:   make(chan error),
		Done:     make(chan struct{}),
	}
}

// ReadLog reads given 'filename' line by line, parse its lines according to
// specified 'Formatter'; can discard results that has timestamp less than 'since'
func (lr *LogsReader) ReadLog(format core.Formatter, since int64) {
	f, err := os.Open(lr.filename)
	if err != nil {
		lr.Errors <- fmt.Errorf("cannot create NewLogFile, %w", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			lr.Errors <- err
			break
		}

		e, err := format(line)
		if err != nil {
			lr.Errors <- fmt.Errorf("cannot format event, %w", err)
		}

		if e.DateTime < since {
			continue
		}

		lr.Events <- e
	}

	lr.Done <- struct{}{}
}
