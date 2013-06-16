package server

import (
	"os/exec"
)

type AptUpdateArgs struct{}

type AptUpdateResults struct {
	Err    string
	Output []byte
}

func (a AptUpdateResults) GetErr() string {
	return a.Err
}

func (a AptUpdateResults) GetOutput() string {
	return string(a.Output)
}

type AptUpdate struct{}

func (t *AptUpdate) Update(args *AptUpdateArgs, results *AptUpdateResults) error {
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
