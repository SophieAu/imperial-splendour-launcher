package backend

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func (a *API) compileSetup() error {

	compileCmd := exec.Command("iscc", a.setupBaseFolder+"/"+setupFile)

	err := compileCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// MAYBE IN THE FUTURE
func handleStdOutAndStdErr() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")

	// get pipes
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// start program
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// handle stdout and stderr
	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	slurpOut, _ := io.ReadAll(stdout)
	fmt.Printf("%s\n", slurpOut)

	// error out if after end of programm something broke
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
