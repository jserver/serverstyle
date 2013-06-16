package server

import (
	"errors"
	"os/exec"
)

type EasyInstallArgs struct {
	Packages []string
}

type EasyInstallResults struct {
	Err    string
	Output []byte
}

func (e EasyInstallResults) GetErr() string {
	return e.Err
}

func (e EasyInstallResults) GetOutput() string {
	return string(e.Output)
}

type EasyInstall struct{}

func (t *EasyInstall) Install(args *EasyInstallArgs, results *EasyInstallResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"easy_install"}
	command = append(command, args.Packages...)
	out, err := exec.Command("sudo", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("Easy Install Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran Easy Install...")
	}
	results.Output = out
	return nil
}
