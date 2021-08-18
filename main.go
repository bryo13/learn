package main

import (
	"github.com/Tum/G-learn/locate"
)

func main() {
	homepath, _ := locate.FindHome()

	locate.MakeDir(homepath)
}
