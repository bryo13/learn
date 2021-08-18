package logger

import (
	"fmt"
	"log"
	"os"
)

type Logs struct {
	l *log.Logger
}

func (lg *Logs) SaveLogs(file *os.File, msg string) error {
	lg.l = log.New(file, msg, log.Llongfile)

	fmt.Println(lg.l)
	return nil
}
