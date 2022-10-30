package linux_agent

import (
	"fmt"

	core "github.com/s-vvardenfell/psyllium/internal"
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
