package main

import (
	"errors"
	"log"
	"os"

	"github.com/fatih/color"
)

const version = "1.0.0"

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()

	if err != nil {
		exitCLI(err)
	}
	setup(arg1, arg2)
	switch arg1 {
	case "help":
		showHelp()
	case "new":
		err := doNew(arg2)
		if err != nil {
			exitCLI(err)
		}
	case "version":
		color.Yellow("version: " + version)
	case "model":
		err := doModel(arg2)
		if err != nil {
			exitCLI(err)
		}
	case "controller":
		err := doController(arg2)
		if err != nil {
			exitCLI(err)
		}
	case "auth":
		err := doAuth()
		if err != nil {
			exitCLI(err)
		}
	default:
		log.Println(arg2, arg3)
		showHelp()
	}
	exitCLI(nil, message)
}

func exitCLI(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if err != nil {
		color.Red("Error: %v\n", err)
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]
		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}
		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error:command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}
	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available Commands:
		help	   - show the help command
		version    - print application version
		model      - create application model
		controller - create application controller
	`)
}
