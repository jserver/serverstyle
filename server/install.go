package server

import (
	"errors"
	"os/exec"
)

type AptInstallArgs struct {
	Packages []string
}

type AptInstallResults struct {
	Err    string
	Output []byte
}

func (a AptInstallResults) GetErr() string {
	return a.Err
}

func (a AptInstallResults) GetOutput() string {
	return string(a.Output)
}

type AptInstall struct{}

func (t *AptInstall) Install(args *AptInstallArgs, results *AptInstallResults) error {
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

