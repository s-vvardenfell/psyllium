package logs_reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	authLog = regexp.MustCompile(
		`([\w]{3,4}\s[\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2}) ([\w|\d|-]{1,40}) (.+)`)

	authLogEvent = regexp.MustCompile(`([\w|\d|_|-]){1,20}:(.+)`)
)

type Event struct {
	DateTime int64
	Host     string
	Process  string
	Msg      string
}

// ReadLog reads given 'filename' line by line, parse its lines according to
// specified 'format'; can discard results that has timestamp less than 'since'
func ReadLog(filename string, format string, since int) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot create NewLogFile, %w", err)
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			break
		}

		// TODO FORMATTER FUNC/INTERFACE
		// parse event
		// parse all res-s to str
		// discard old ones
		match := authLog.FindStringSubmatch(line)
		if match == nil {
			return fmt.Errorf("failed to parse log str <%s> by FindAllString", line) //todo json-log
		}

		dt := match[1] // todo check
		host := match[2]
		data := match[3]

		match = authLogEvent.FindStringSubmatch(line)
		if match == nil {
			return fmt.Errorf("failed to parse log str <%s> by FindAllString", line) //todo json-log
		}

		udt, err := time.Parse("Jan 02 15:04:05", dt)
		if err != nil {
			logrus.Errorf("cannot parse time, %v", err) //todo send err to chan
		}
		udt = udt.AddDate(time.Now().Year(), 0, 0)

		e := Event{
			DateTime: udt.Unix(),
			Host:     host,
			Process:  match[1],
			Msg:      match[2],
		}

		fmt.Printf("%#v\n\n", e)

	}

	return nil
}
