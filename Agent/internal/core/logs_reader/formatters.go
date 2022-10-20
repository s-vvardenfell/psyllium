package logs_reader

import (
	"fmt"
	"strings"
	"time"
)

type Formatter func(event string) (*Event, error)

func FormatSysLog(event string) (*Event, error) {
	match := authLog.FindStringSubmatch(event)
	if match == nil || len(match) < 3 {
		return nil, fmt.Errorf("failed to process log event <%s>", event) //todo json-log
	}

	var (
		dt   = match[1]
		host = match[2]
		data = match[3]
	)

	udt, err := time.Parse("Jan 02 15:04:05", dt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse event time, %w", err)
	}

	udt = udt.AddDate(time.Now().Year(), 0, 0)

	events := strings.Split(data, ": ")

	return &Event{
		DateTime: udt.Unix(),
		Host:     host,
		Process:  events[0],
		Msg:      strings.Join(events[1:], " "),
	}, nil
}
