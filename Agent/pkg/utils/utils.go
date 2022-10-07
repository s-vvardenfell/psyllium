package utils

import (
	"agent/internal/core"
	"bufio"
	"context"
	"io"
	"os"
	"time"
)

/*
// usage example:

	filename := "file"
	dataChan := make(chan string)
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go ReadFileToChan(ctx, filename, dataChan, errChan)

	go func() {
		time.Sleep(30 * time.Second)
		cancel()
	}()

	for {
		select {
		case data := <-dataChan:
			fmt.Println(data)
		case err := <-errChan:
			fmt.Printf("Selected err case, %v\n", err)
			return
		}
	}
*/

// ReadFileToChan opens a given file by its name and reads it line by line
// to channel 'dataCh' with frequency 'freq'; Writes error to 'errCh' if occurs
func ReadFileToChan(
	ctx context.Context, filename string, dataCh chan<- core.Event, errCh chan error, freq int) {

	f, err := os.Open(filename)
	if err != nil {
		errCh <- err
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
		case <-ticker.C:
			line, err := rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					rd.Reset(f)
				} else {
					errCh <- err
					return
				}
			}

			dataCh <- core.Event{What: line}
		}
	}
}
