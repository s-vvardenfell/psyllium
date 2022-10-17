package main

import (
	"fmt"
)

type HostInfo struct {
	OS       string
	Host     string `env:"HOST"`
	Home     string `env:"HOME"`
	Username string `env:"USERNAME"`
	Shell    string `env:"SHELL"`
	Term     string `env:"TERM"`
}

func main() {
	cfg := HostInfo{}
	// if err := env.Parse(&cfg); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }

	fmt.Printf("%+v\n", cfg)

	// cmd.Execute()

	// logs := []string{"test/file1.log", "test/file2.log"}

	// lr, err := logs_reader.NewLogsReader(logs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res := lr.Work()

	// for r := range res {
	// 	fmt.Println(r)
	// }

	// 	l1, err := logs_reader.NewLogFile(logs[0])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	l2, err := logs_reader.NewLogFile(logs[1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	events := make(chan string)
	// 	errs := make(chan error)
	// 	done := make(chan struct{})

	// 	go func() {
	// 		time.Sleep(3 * time.Second)
	// 		l1.ReadOldEvents(events, errs, done)
	// 	}()

	// 	go func() {
	// 		time.Sleep(3 * time.Second)
	// 		go l2.ReadOldEvents(events, errs, done)
	// 	}()

	// 	cnt := len(logs)

	// 	for cnt != 0 {
	// 		select {
	// 		case event := <-events:
	// 			fmt.Println(event)
	// 		case err := <-errs:
	// 			fmt.Println(err)
	// 		case <-done:
	// 			cnt--
	// 		}
	// 	}

	// 	fmt.Println("ENDED READING OLD EVENTS")

	// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// 	defer cancel()

	// 	go func() {
	// 		time.Sleep(2 * time.Second)
	// 		l1.ReadNewEvents(ctx, events, errs, 5)
	// 	}()

	// 	go func() {
	// 		time.Sleep(2 * time.Second)
	// 		go l2.ReadNewEvents(ctx, events, errs, 5)
	// 	}()

	// loop:
	// 	for {
	// 		select {
	// 		case event := <-events:
	// 			fmt.Println(event)
	// 		case err := <-errs:
	// 			fmt.Println(err)
	// 			break loop
	// 		}
	// 	}

	// fmt.Println("Breaked loop")
}

func NewLogsReader(logs []string) {
	panic("unimplemented")
}
