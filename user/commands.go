package user

import (
	"errors"
	"fmt"
)

type Command struct {
	Name          string
	Function      func(args ...[]string) error
	Children      []Command
	FunctionNames []string
	Functions     []func(args ...[]string) error
}

func (command Command) Execute() (string, error) {
	// ideally there would be a value added to stored variables
	return "", nil
}

func printHelp(args ...[]string) error {
	if args != nil {
		return errors.New("ArgumentError: 'help' does not accept arguments")
	}

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

	return nil
}

var (
	ShowHelp = Command{Name: "help", Function: printHelp}
)
