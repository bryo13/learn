package locate

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

// FindHome return a string the location of home or an error using homedir package
func FindHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", errors.New("Could not find home dir")
	}

	return home, nil
}

// first return string is dir locationand the second is the name plus location
func specifyLocation(path, nm string) (string, string) {
	locat, _ := FindHome()
	path = fmt.Sprintf("%s/%s", locat, "glearn")
	fileToUse := fmt.Sprintf("%s/%s", path, nm)

	return path, fileToUse
}

// MakeDir take the location of the home dir and makes a folder if none exists or returns an error
func MakeDir(home string) (bool, error) {
	home, _ = FindHome()
	homedir, err := os.ReadDir(home)
	dirName, _ := specifyLocation(home, "glearn.log")
	if err != nil {
		return false, errors.New("Could not read home dir")
	}

	for _, name := range homedir {
		if name.Name() != "glearn" {
			os.Mkdir(dirName, 0755)
			fmt.Println("Created ")
			return false, nil
		}
		break
	}
	return true, nil
}
