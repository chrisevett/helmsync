package util

import (
	"bytes"
	"log"
	"os/exec"
)

var execCommand = exec.Command

func Command(cmd *exec.Cmd) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	returnText := ""
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			returnText = stderr.String()
			log.Printf("Error: %s", returnText)
		}
	} else {
		returnText = stdout.String()
		log.Printf("Success: %s", returnText)
	}
	return returnText, err
}
