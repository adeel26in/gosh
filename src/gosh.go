package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	useros := runtime.GOOS

	if useros != "windows" {
		reader := bufio.NewReader(os.Stdin)
		for {

			cwd, err := os.Getwd()
			if err != nil {
				log.Println("Can't get the directory!")
			}

			fmt.Print(cwd, "> ")
			// Read the keyboad input.
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			// Handle the execution of the input.
			if err = execInputNONWINDOWS(input); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}

	if useros == "windows" {
		reader := bufio.NewReader(os.Stdin)
		for {

			cwd, err := os.Getwd()
			if err != nil {
				log.Println("Can't get the directory!")
			}

			fmt.Print(cwd, "> ")
			// Read the keyboad input.
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			// Handle the execution of the input.
			if err = execInputWindows(input); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
