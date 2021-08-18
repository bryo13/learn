package logger

import (
	"fmt"
	"log"
	"os"
)

// Logs defines a log
type Logs struct {
	l *log.Logger
}

// SaveLogs takes takes an instance of an open file and a message and returns an error
func (lg *Logs) SaveLogs(file *os.File, msg string) error {
	lg.l = log.New(file, msg, log.Llongfile)

	fmt.Println(lg.l)
	return nil
}
