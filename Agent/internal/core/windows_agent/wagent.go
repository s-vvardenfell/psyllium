package windows_agent

import (
	"agent/internal/core"
	"fmt"
)

type WindowsAgent struct {
	msg string
}

func NewWindowsAgent() *WindowsAgent {
	return &WindowsAgent{
		msg: "WindowsAgent works!",
	}
}

func (l *WindowsAgent) GetEvents() (c <-chan core.Events) {
	fmt.Println(l.msg)
	return
}
