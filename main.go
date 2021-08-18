package main

import (
	"os"

	"github.com/Tum/G-learn/locate"
	"github.com/Tum/G-learn/logger"
)

func main() {
	lf := new(locate.Files)
	lgs := new(logger.Logs)

	homepath, _ := lf.FindHome()
	lgs.SaveLogs(os.Stdout, homepath)
	lf.MakeDir(homepath)
}
