package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
		//return nil, errors.WithMessage(err, "open failed")
		//return nil, errors.WithStack(err)
	}

	defer f.Close()
	return nil, nil
}

func ReadConfig() ([]byte, error) {

	home := os.Getenv("Home")
	config, err := ReadFile(filepath.Join(home, "xxx"))
	//return config, errors.WithMessage(err, "could not read config")
	return config, errors.Wrap(err, "could not read config")
}

func main() {

	if _, err := ReadFile("test"); err != nil {

		fmt.Printf("ReadFile: %v\n\n", errors.Cause(err))
		fmt.Printf("ReadFilestack trace: %+v\n\n", err)
		fmt.Println(err)
		fmt.Println()
	}

	if _, err := ReadConfig(); err != nil {
		fmt.Printf("ReadConfig: %v\n\n", errors.Cause(err))
		fmt.Printf("ReadConfig stack2 trace: %+v\n\n", err)
		fmt.Println(err)
	}
}
