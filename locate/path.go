package locate

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

type Files interface {
	FindHome() (string, error)
	MakeDir(string) (bool, error)
	OpenFile(string) (*os.File, error)
	SpecifyLocation(string, string) (string, string)
}

type LocateFiles struct{}

// FindHome return a string the location of home or an error using homedir package
func (l *LocateFiles) FindHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", errors.New("Could not find home dir")
	}

	return home, nil
}

// first return string is dir location and the second is the file name plus location
func (l *LocateFiles) SpecifyLocation(folderName, path, nm string) (string, string) {
	locat, _ := l.FindHome()
	path = fmt.Sprintf("%s/%s", locat, folderName)
	fileToUse := fmt.Sprintf("%s/%s", path, nm)

	return path, fileToUse
}

// MakeDir take the location of the home dir and makes a folder if none exists or returns an error
func (l *LocateFiles) MakeDir(home, folderName string) (bool, error) {
	home, _ = l.FindHome()
	homedir, err := os.ReadDir(home)
	dirName, _ := l.SpecifyLocation(folderName, home, "glearn.log")
	if err != nil {
		return false, errors.New("Could not read home dir")
	}

	for _, name := range homedir {
		if name.Name() != folderName {
			os.Mkdir(dirName, 0755)
			fmt.Println("Created ")
			return false, nil
		}
		break
	}
	return true, nil
}

func (l *LocateFiles) OpenFile(path string) {
	// To open we need location to file
}
