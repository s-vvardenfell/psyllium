package linux_agent

import (
	"agent/internal/core"
	"agent/pkg/utils"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	chanCap = 10
	freq    = 5
)

type LogFile struct {
	File     *os.File
	FileName string
	LastLine int // Last line read
}

type LinuxAgent struct {
	msg      string // TEST
	events   chan<- core.Event
	errors   chan error
	logFiles []LogFile
}

func NewLinuxAgent(logfiles []string) core.Agent {
	if len(logfiles) == 0 {
		logrus.Errorf("logfiles list is empty")
		return nil
	}

	lgf := make([]LogFile, 0, len(logfiles))

	for i := range logfiles {
		lgf = append(lgf, LogFile{
			FileName: logfiles[i],
		})
	}

	return &LinuxAgent{
		msg:      "LinuxAgent works!",
		events:   make(chan<- core.Event, chanCap),
		errors:   make(chan error),
		logFiles: lgf,
	}
}

func readLogFile() {

}

func (l *LinuxAgent) GetEvents() {
	fmt.Println(l.msg)

	for _, file := range l.logFiles {
		go utils.ReadFileToChan(context.Background(), file.FileName, l.events, l.errors, freq)
	}

	return
}

func (l *LinuxAgent) Err() <-chan error {
	return l.errors
}
