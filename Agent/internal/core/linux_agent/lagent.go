package linux_agent

import (
	"agent/internal/core"
	"fmt"
)

type LinuxAgent struct {
	msg string // TEST

}

func NewLinuxAgent() core.Agent {
	return &LinuxAgent{
		msg: "LinuxAgent works!",
	}
}

func (l *LinuxAgent) Run() {

}

func (l *LinuxAgent) GetEvents() {
	fmt.Println(l.msg)

	// return
}
