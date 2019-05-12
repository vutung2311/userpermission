package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"userpermission/internal/usecase"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmdStrings := make([]string, 0, 1000)
	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("Got error %v while reading from stdin", err)
		}
		if cmdString = strings.Trim(cmdString, "\n\t"); len(cmdString) > 0 {
			cmdStrings = append(cmdStrings, cmdString)
		}
		if err == io.EOF {
			break
		}
	}
	output, err := usecase.ProcessPermissionInput(cmdStrings)
	if err != nil {
		log.Fatalf("Got error %v while processing input", err)
	}
	fmt.Println(strings.Join(output, "\n"))
}
