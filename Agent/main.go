package main

import (
	"agent/internal/core/logs_reader"
	"fmt"
)

var data = "gdm-launch-environment]: pam_unix(gdm-launch-environment:session): session opened for user gdm(uid=127) by (uid=0)"

func main() {
	// fmt.Println("works!")

	lm := logs_reader.NewLogReader("test/auth.log")
	go lm.ReadLog(logs_reader.FormatSysLog, 0)

loop:
	for {
		select {
		case ev := <-lm.Events:
			fmt.Println(ev)
		case err := <-lm.Errors:
			fmt.Println("Error!")
			fmt.Println(err)
		case <-lm.Done:
			fmt.Print("DONE!")
			break loop
		}
	}

	///////////////////////

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
