package main

import (
	"agent/internal/core"
	"agent/internal/core/logs_reader"
	"fmt"
	"log"
)

var files = []string{"test/file1.log", "test/file2.log"}

func main() {
	fmt.Println("works!")
	lfr, err := logs_reader.NewLogsReader(files)
	if err != nil {
		log.Fatal(err)
	}

	events := make(chan core.Event)
	errors := make(chan error)
	done := make(chan struct{})

	lfr.Work(events, done, errors)

	l := len(files)

loop:
	for l != 0 {
		select {
		case ev := <-events:
			fmt.Println(ev)
		case err := <-errors:
			fmt.Println(err)
			break loop
		case <-done:
			l--
			fmt.Println("DONE")
		}
	}

}
