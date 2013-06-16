package server

import (
	"os/exec"
)

type AptGetUpdateArgs struct {}

type AptGetUpdateResults struct {
	Err    string
	Output []byte
}

func (a AptGetUpdateResults) GetErr() string {
	return a.Err
}

func (a AptGetUpdateResults) GetOutput() string {
	return string(a.Output)
}

type AptGetUpdate struct {}

func (t *AptGetUpdate) Update(args *AptGetUpdateArgs, results *AptGetUpdateResults) error {
	command := []string{"apt-get", "-y", "update"}
	out, err := exec.Command("sudo", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Update Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Update...")
	}
	results.Output = out
	return nil
}
