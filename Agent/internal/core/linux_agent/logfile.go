package linux_agent

import (
	"bufio"
	"fmt"
	"os"
)

// TODO other pkg or core lvl
type LogFile struct {
	File     *os.File
	FileName string
	LastLine int // Last line read
}

// wg to args?
func (l *LogFile) readOldEvents(events chan<- string, errs chan<- error, done chan<- int) {
	var err error
	fmt.Printf("START READ %s\n", l.FileName) // TEST
	l.File, err = os.Open(l.FileName)
	if err != nil {
		errs <- err
		return
	}

	sc := bufio.NewScanner(l.File)
	cnt := 1 // TEST
	for sc.Scan() {
		events <- fmt.Sprintf("# %d %s file: %s\n", cnt, sc.Text(), l.FileName)
		cnt++
	}

	if err := sc.Err(); err != nil {
		errs <- fmt.Errorf("scan file error: %v", err)
		return
	}

	done <- 1
}

func (l *LogFile) readNewEvents() {

}

func (l *LogFile) CleanUp() {
	_ = l.File.Close()
}
