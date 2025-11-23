package main

import (
	"os"
	"os/exec"
	"strings"
)

func execInputNONWINDOWS(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Fields(input)

	if len(args) == 0 {
		return nil
	}

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir(os.Getenv("HOME"))
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command and return the error.
	return cmd.Run()
}

func execInputWindows(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Fields(input)

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir(os.Getenv("USERPROFILE"))
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	// For other commands, execute them directly
	cmd := exec.Command("cmd", "/C", strings.Join(args, " "))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	return cmd.Run()
}
