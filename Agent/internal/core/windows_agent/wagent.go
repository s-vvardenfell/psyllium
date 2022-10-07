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

func (l *WindowsAgent) GetEvents() {
	fmt.Println(l.msg)
	return
}
