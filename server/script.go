package server

import (
	"errors"
	"os"
	"os/exec"
)

type ScriptArgs struct {
	Name    string
	Content []byte
}

type ScriptResults struct {
	Err    string
	Output []byte
}

func (s ScriptResults) GetErr() string {
	return s.Err
}

func (s ScriptResults) GetOutput() string {
	return string(s.Output)
}

type Script struct{}

func (t *Script) Runner(args *ScriptArgs, results *ScriptResults) error {
	if len(args.Name) == 0 || len(args.Content) == 0 {
		return errors.New("no script to run")
	}

	script := "/var/lib/serverstyle/scripts/" + args.Name

	file, err := os.Create(script)
	if err != nil {
		return errors.New("unable to create script file")
	}
	_, err = file.Write(args.Content)
	if err != nil {
		return errors.New("unable to write script file")
	}
	err = file.Close()
	if err != nil {
		return errors.New("unable to close script file")
	}
	err = os.Chmod(script, 0755)
	if err != nil {
		return errors.New("unable to chmod script file")
	}

	out, err := exec.Command(script).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("Script Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran Script...")
	}
	results.Output = out
	_ = os.Remove(script)
	return nil
}
