package server

import (
	"errors"
	"os/exec"
)

type AptGetInstallArgs struct {
	Packages []string
}

type AptGetInstallResults struct {
	Err    string
	Output []byte
}

func (a AptGetInstallResults) GetErr() string {
	return a.Err
}

func (a AptGetInstallResults) GetOutput() string {
	return string(a.Output)
}

type AptGetInstall struct {}

func (t *AptGetInstall) Install(args *AptGetInstallArgs, results *AptGetInstallResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"apt-get", "-y", "install"}
	command = append(command, args.Packages...)
	out, err := exec.Command("sudo", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Install Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Install...")
	}
	results.Output = out
	return nil
}
