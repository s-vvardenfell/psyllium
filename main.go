package main

import (
	"fmt"
	"log"

	"github.com/s-vvardenfell/psyllium/internal/checker"
	"github.com/s-vvardenfell/psyllium/internal/checker/shhchecker"
)

func main() {
	// cmd.Execute()

	s := shhchecker.New(checker.CheckOptions{
		CheckType: "ssh",
		FileNames: []string{"known_hosts", "authorized_keys"},
	})

	res := s.Check()

	if res.Err != nil {
		log.Fatal(res.Err)
	}

	fmt.Printf("%#v", res)
}
