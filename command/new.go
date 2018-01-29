package command

import (
	"fmt"
	"os"
	"os/user"

	"github.com/urfave/cli"
)

var (
	gotDirPath = ""
)

func init() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	gotDirPath = user.HomeDir + "/.got"
}

func New(c *cli.Context) error {
	ok, err := Exists(gotDirPath)
	if err != nil {
		return err
	}
	if ok {
		err = downloadTemplate()
		if err != nil {
			return err
		}
	} else {
		err = first()
		if err != nil {
			return err
		}
		err = downloadTemplate()
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadTemplate() error {
	fmt.Println("downloading...")
	return nil
}

func first() error {
	err := os.Mkdir(gotDirPath, 0777)
	if err != nil {
		return err
	}
	return nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
