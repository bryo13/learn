package main

import (
	"os"

	"github.com/Tum/G-learn/locate"
	"github.com/Tum/G-learn/logger"
)

func main() {
	lf := new(locate.LocateFiles)
	lgs := new(logger.Logs)

	homepath, _ := lf.FindHome()
	lf.MakeDir(homepath, "glearn")
	lgs.SaveLogs(homepath, os.Stdout, homepath)
}
