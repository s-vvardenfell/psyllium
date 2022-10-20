package main

import (
	"agent/internal/core/logs_reader"
	"fmt"

	"github.com/sirupsen/logrus"
)

var files = []string{"test/file1.log", "test/file2.log"}

func main() {
	fmt.Println("works!")

	err := logs_reader.ReadLog("test/auth.log", "", 100)
	if err != nil {
		logrus.Error(err)
	}

	// 	lfr, err := logs_reader.NewLogsReader(files)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	events := make(chan core.Event)
	// 	errors := make(chan error)
	// 	done := make(chan struct{})

	// 	lfr.Work(events, done, errors)

	// 	l := len(files)

	// loop:
	//
	//	for l != 0 {
	//		select {
	//		case ev := <-events:
	//			fmt.Println(ev)
	//		case err := <-errors:
	//			fmt.Println(err)
	//			break loop
	//		case <-done:
	//			l--
	//			fmt.Println("DONE")
	//		}
	//	}
}
