package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// timeout after 1 minute
var TimeoutSeconds int = 60

var CLICommands = []Command{
	ShowHelp,
}

func CommandLineInterface() {
	fmt.Println("Welcome to Particle Simulator CLI! type 'help' for options.")

	var commandString string
	var err error = nil

	startTime := time.Now()

	for {
		// exit command
		if commandString == "exit" {
			break
		}

		secondsPassed := time.Since(startTime).Seconds()

		if secondsPassed > float64(TimeoutSeconds) {
			break
		}

		// Wait for user to input command or timeout
		commandString, err = GetCommand(TimeoutSeconds)

		if err != nil {
			fmt.Println(err)
		} else {
			for i := len(CLICommands) - 1; i >= 0; i-- {
				result, err := CLICommands[i].Execute()

				// ignore errors, they indicate if the command can execute the given command
				if err != nil {
					// display result string
					fmt.Println(result)
				}
			}
		}
	}
}

func GetCommand(timeoutSeconds int) (string, error) {

	inputChannel := make(chan string)
	errorChannel := make(chan error)

	go func() {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		// Open input to user and get a line
		if scanner.Scan() {
			// Send the line read through the input channel after trimming white space
			command := strings.TrimSpace(scanner.Text())
			inputChannel <- command
		} else {
			// If there is an error instead send it to the error channel
			errorChannel <- scanner.Err()
		}
	}()

	// Wait and listen on all three channels for a timeout, input or an error, then return based on that
	select {
	case input := <-inputChannel:
		return input, nil
	case err := <-errorChannel:
		return "", err
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		// Timout occured first
		err := fmt.Errorf("\ntimed out after %d seconds", timeoutSeconds)
		return "", err
	}
}
