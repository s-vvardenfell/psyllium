package host

import (
	"fmt"
	"runtime"

	"github.com/caarlos0/env"
)

type HostInfo struct {
	OS       string
	Host     string `env:"HOST"`
	Home     string `env:"HOME"`
	Username string `env:"USERNAME"`
	Shell    string `env:"SHELL"`
	Term     string `env:"TERM"`
}

// TODO make drossel or other pattern to add pause / return cached result
func GetHostInfo() *HostInfo {
	inf := HostInfo{}
	if err := env.Parse(&inf); err != nil {
		fmt.Printf("%+v\n", err)
	}

	inf.OS = runtime.GOOS

	return &inf
}
