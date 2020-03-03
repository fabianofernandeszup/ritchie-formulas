package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os/exec"
	"ritchie/pkg/ritchie/util/error"
	"ritchie/pkg/ritchie/util/input"
	"ritchie/pkg/ritchie/util/makefile"
	"ritchie/pkg/ritchie/util/path"
	"ritchie/pkg/ritchie/util/template"
	"ritchie/pkg/ritchie/util/tree"
	"strings"
)

func main() {

	inputValue := input.BuildInput()
	mainPaths := path.BuildMainPaths()
	if !path.IsOnRightDir(mainPaths) {
		return
	}

	tree.ChangeTreeFile(inputValue, mainPaths)
	template.GenerateFiles(inputValue, mainPaths, 0)
	makefile.ChangeMakeFile(inputValue, mainPaths)

	execCommand("make test-local form=" + strings.ToUpper(inputValue.Name))

	color.Green("Generate formula:" + inputValue.Name + " with description:" + inputValue.Description + " .")

}

func execCommand(value string) string {
	command := strings.Split(value, " ")[0]
	params := strings.Split(value, " ")[1:]
	log.Printf("Executing command: %v params: %v\n", command, params)
	cmd := exec.Command(command, params...)
	stdout, _ := cmd.StdoutPipe()
	var outError bytes.Buffer
	cmd.Stderr = &outError
	error.VerifyError(cmd.Start())
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	commandResultMessage := ""
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		commandResultMessage += m
	}
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("Failed to execute command %v\nParams: %v\nError: %v", command, params, outError.String())
	}
	return commandResultMessage
}
