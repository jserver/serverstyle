package server

import (
	"errors"
	"os/exec"
)

type AptGetArgs struct {
	Packages []string
}

type AptGetResults struct {
	Output []byte
	Err string
}

type AptGet int

func (t *AptGet) Install(args *AptGetArgs, results *AptGetResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"apt-get", "-y", "install"}
	command = append(command, args.Packages...)
	out, err := exec.Command("sudo", command...).CombinedOutput()
	results.Err = err.Error()
	results.Output = out
	return nil
}
