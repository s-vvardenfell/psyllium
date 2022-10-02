package linux_agent

import (
	"agent/internal/core"
	"fmt"
)

type LinuxAgent struct {
	msg string
}

func NewLinuxAgent() *LinuxAgent {
	return &LinuxAgent{
		msg: "LinuxAgent works!",
	}
}

func (l *LinuxAgent) GetEvents() (c <-chan core.Events) {
	fmt.Println(l.msg)
	return
}
