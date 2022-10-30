package core

import "context"

type Effector func(context.Context) (*HostInfo, error)

type Formatter func(event string) (*Event, error)

type HostInfo struct {
	OS       string `json:"os"`
	Host     string `env:"HOSTNAME" json:"host"`
	Home     string `env:"HOME" json:"home"`
	Username string `env:"USERNAME" json:"uname"`
	Shell    string `env:"SHELL" json:"shell"`
	Term     string `env:"TERM" json:"term"`
}

type Event struct {
	DateTime int64  `json:"event_dt"`
	Host     string `json:"host"`
	Process  string `json:"process"`
	Msg      string `json:"msg"`
}

type Msg struct {
	HostInfo HostInfo `json:"hostinfo"`
	Events   []Event  `json:"events"`
}
