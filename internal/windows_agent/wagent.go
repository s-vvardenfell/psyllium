package windows_agent

import (
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

func (l *WindowsAgent) Run() {
	fmt.Println(l.msg)
	// return
}
