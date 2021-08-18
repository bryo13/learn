package main

import "github.com/Tum/G-learn/locate"

func main() {
	lf := new(locate.LocateFiles)
	homepath, _ := lf.FindHome()

	lf.MakeDir(homepath)
}
