package utils

import (
	"bufio"
	"context"
	"io"
	"os"
	"time"
)

/*
	filename := "file"
	dataChan := make(chan string)
	errChan := make(chan error)

	// go ReadFile(filename, dataChan, errChan)
	go ReadFile2(filename, dataChan, errChan)

	for {
		select {
		case data := <-dataChan:
			fmt.Println(data)
		case err := <-errChan:
			fmt.Printf("Selected err case, %v\n", err)
			// default:
			// 	continue
			// 	break
		}
	}
*/

/*
TODO
use ctx
args to struct
сравнить с вариантом из "облачного го"
сделать селект, в одном кейсе будет тикер, в другом контекст на отмену - попробовать в соседней функции-копии
*/
func ReadFileToChan(ctx context.Context, filename string, dataCh chan<- string, errCh chan<- error) {
	f, err := os.Open(filename)
	if err != nil {
		errCh <- err
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				rd.Reset(f)
			} else {
				errCh <- err
				return
			}

			time.Sleep(3 * time.Second)
		}

		dataCh <- line
	}
}
