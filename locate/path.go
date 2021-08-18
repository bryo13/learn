package locate

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

// Locate interface is used to define a file and dir location
type Locate interface {
	FindHome() (string, error)
	OpenFileInGlearn(path string) (*os.File, error)
	MakeDir(string) (bool, error)
	SpecifyLocation(string, string) (string, string)
}

// Files used to initialize Files interface
type Files struct{}

// FindHome return a string the location of home or an error using homedir package
func (f *Files) FindHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", errors.New("Could not find home dir")
	}

	return home, nil
}

// SpecifyLocation irst return string is dir location and the second is the file name plus location
func (f *Files) SpecifyLocation(path, nm string) (string, string) {
	locat, _ := f.FindHome()
	path = fmt.Sprintf("%s/%s", locat, "glearn")
	fileToUse := fmt.Sprintf("%s/%s", path, nm)

	return path, fileToUse
}

// MakeDir take the location of the home dir and makes a folder if none exists or returns an error
func (f *Files) MakeDir(home string) (bool, error) {
	home, _ = f.FindHome()
	homedir, err := os.ReadDir(home)
	dirName, _ := f.SpecifyLocation(home, "glearn.log")
	if err != nil {
		return false, errors.New("Could not read home dir")
	}

	// Os.ReadDir() returns a slice of DirEntry type hence one loop through them. A name function is availed in the DirEntry interface
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
