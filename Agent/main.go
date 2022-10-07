package main

import (
	"agent/internal/core"
	"agent/internal/core/linux_agent"
	"agent/internal/frontend"
	"log"
)

func main() {
	// cmd.Execute()
	logs := []string{"/var/log/auth.log", "/var/log/auth.log1"}
	la := linux_agent.NewLinuxAgent(logs)
	// wa := windows_agent.NewWindowsAgent()

	c := core.NewCore()
	c.WithAgent(la)

	front, err := frontend.NewFrontEnd("rest")
	if err != nil {
		log.Fatal(err)
	}

	front.StartWith(c)
}
