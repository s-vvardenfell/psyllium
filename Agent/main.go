package main

import (
	"fmt"
	"go/types"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// fmt.Println("works!")

	cnfg := types.Config{}
	viper.SetConfigFile("configs/config.yml")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		cobra.CheckErr(err)
	}

	if err := viper.Unmarshal(&cnfg); err != nil {
		cobra.CheckErr(err)
	}

	fmt.Println(cnfg)

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
