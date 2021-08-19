package logger

import (
	"fmt"
	"log"
	"os"
)

<<<<<<< HEAD
// Logs defines a log
=======
// Logs defines a l of type *Log.logger
>>>>>>> logger
type Logs struct {
	l *log.Logger
}

<<<<<<< HEAD
// SaveLogs takes takes an instance of an open file and a message and returns an error
func (lg *Logs) SaveLogs(file *os.File, msg string) error {
=======
// SaveLogs take the defination of Logs and saves it in a file
func (lg *Logs) SaveLogs(path string, file *os.File, msg string) error {
	// variables
>>>>>>> logger
	lg.l = log.New(file, msg, log.Llongfile)

	// Open folder to save logs
	// Open file
	// set location
	// write log file
	// defer close file
	fmt.Println(lg.l)
	return nil
}
