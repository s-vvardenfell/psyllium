package logs_reader

import "errors"

var (
	ErrLogFileListEmpty = errors.New("log file list is empty")
	chanCap             = 0
)

type LogFileReader struct {
	events   chan string
	errors   chan error
	done     chan int
	logFiles []LogFile
}

func NewLogsReader(logfiles []string) (*LogFileReader, error) {
	// if len(logfiles) == 0 {
	// 	return nil, ErrLogFileListEmpty
	// }

	// lgf := make([]LogFile, 0, len(logfiles))

	// for i := range logfiles {
	// 	lgf = append(lgf, LogFile{
	// 		FileName: logfiles[i],
	// 	})
	// }

	// return &LogFileReader{
	// 	// events:   make(chan<- core.Event, chanCap),
	// 	events:   make(chan string, chanCap),
	// 	errors:   make(chan error),
	// 	done:     make(chan int),
	// 	logFiles: lgf,
	// }, nil
	return &LogFileReader{}, nil
}
