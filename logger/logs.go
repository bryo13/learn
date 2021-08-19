package logger

import (
	"fmt"
	"log"
	"os"
)

// Logs defines a l of type *Log.logger
type Logs struct {
	l *log.Logger
}

// SaveLogs take the defination of Logs and saves it in a file
func (lg *Logs) SaveLogs(path string, file *os.File, msg string) error {
	// variables
	lg.l = log.New(file, msg, log.Llongfile)

	// Open folder to save logs
	// Open file
	// set location
	// write log file
	// defer close file
	fmt.Println(lg.l)
	return nil
}
