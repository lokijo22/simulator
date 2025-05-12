package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Particle Simulator CLI! type 'help' for options.")

	var fails int = 0

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // EOF or Error
		}

		input := scanner.Text() // Read user input
		command := strings.TrimSpace(input)

		switch {
		case command == "help":
			fails = 0
			printHelp()

		case command == "exit":
			fails = 0
			fmt.Println("Exiting...")
			return

		default:
			if fails == 5 {
				fmt.Println("Too many failed attempts. Exiting...")
				return
			}

			fmt.Println("Unrecognized command")
			fails++
		}
	}
}

func printHelp() {
	help_str := "All these steps must be completed before a simulation can be ran:\n"
	help_str += "  create simulation -name\n"
	help_str += "  create environment\n"
	help_str += "  create space -space_type\n"
	help_str += "  create particles --random --empty\n"
	help_str += "  space add particles\n"
	help_str += "  environment add space\n"
	help_str += "  simulation add environment\n\n"
	help_str += "  create rule -name"
	help_str += "  simulation add rule"
	fmt.Println(help_str)
}
