package main

import (
	"agent/internal/core"
	"agent/internal/core/host_info"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// fmt.Println("works!")

	h, err := host_info.GetHostInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v\n", h)

	e1 := core.Event{
		DateTime: 12345,
		Host:     "localhost",
		Process:  "angryMalvare",
		Msg:      "preparing to encryption...",
	}

	e2 := core.Event{
		DateTime: 12346,
		Host:     "localhost",
		Process:  "angryMalvare",
		Msg:      "start encryption...",
	}

	m := core.Msg{
		HostInfo: h,
		Events:   []core.Event{e1, e2},
	}

	res, err := json.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))

	// 	lm := logs_reader.NewLogsReader("test/auth.log")
	// 	go lm.ReadLog(logs_reader.FormatSysLog, 0)

	// loop:
	// 	for {
	// 		select {
	// 		case ev := <-lm.Events:
	// 			fmt.Println(ev)
	// 		case err := <-lm.Errors:
	// 			fmt.Println("Error!")
	// 			fmt.Println(err)
	// 		case <-lm.Done:
	// 			fmt.Print("DONE!")
	// 			break loop
	// 		}
	// 	}

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
