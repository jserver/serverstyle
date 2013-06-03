package server

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type ScriptArgs struct {
	Name string
	Content string
}

type ScriptResults struct {
	Output []byte
	Err string
}

type Script int

func (t *Script) Runner(args *ScriptArgs, results *ScriptResults) error {
	if len(args.Name) == 0 || len(args.Content) == 0 {
		return errors.New("no script to run")
	}

	file, err := os.Create(args.Name)
	if err != nil {
		return errors.New("unable to create script file")
	}
	_, err = file.WriteString(args.Content)
	if err != nil {
		return errors.New("unable to write script file")
	}
	err = file.Close()
	if err != nil {
		return errors.New("unable to close script file")
	}
	err = file.Chmod(0755)
	if err != nil {
		return errors.New("unable to chmod script file")
	}

	command := "./" + args.Name
	out, err := exec.Command(command).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		fmt.Println("Script Error [ " + results.Err + " ]")
	} else {
		fmt.Println("Successfully ran Script...")
	}
	results.Output = out
	return nil
}
