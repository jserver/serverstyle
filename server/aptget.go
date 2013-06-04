package server

import (
	"errors"
	"fmt"
	"os/exec"
)

type AptGetArgs struct {
	Packages []string
}

type AptGetResults struct {
	Err string
	Output []byte
}

func (a AptGetResults) GetErr() string {
	return a.Err
}

func (a AptGetResults) GetOutput() string {
	return string(a.Output)
}

type AptGet int

func (t *AptGet) Install(args *AptGetArgs, results *AptGetResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"apt-get", "-y", "install"}
	command = append(command, args.Packages...)
	out, err := exec.Command("sudo", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		fmt.Println("AptGet Error [ " + results.Err + " ]")
	} else {
		fmt.Println("Successfully ran AptGet...")
	}
	results.Output = out
	return nil
}
