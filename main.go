package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/projectdiscovery/gologger"
	"golang.org/x/term"
	// Assuming gologger is properly imported.
)

func main() {
	fmt.Println("int(os.Stdin.Fd()): ", int(os.Stdin.Fd()))
	var apiKeyBytes []byte
	var err error
	if runtime.GOOS == "windows" {
		apiKeyBytes, err = readPasswordFromWindows("[*] Enter PDCP API Key (exit to abort): ")
	} else {
		apiKeyBytes, err = readPasswordFromUnix("[*] Enter PDCP API Key (exit to abort): ")
	}
	if err != nil {
		gologger.Fatal().Msgf("Could not read input from terminal: %s\n", err)
	}
	apiKey := string(apiKeyBytes)

	fmt.Println("\nAPI Key entered:", apiKey) // Be cautious with printing sensitive info like API keys.
	if apiKey == "exit" {
		fmt.Println("Exiting.")
		os.Exit(0)
	}

	// Proceed with using the apiKey...
}

func readPasswordFromUnix(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	tty, err := os.Open("/dev/tty")
	if err != nil {
		return nil, err
	}
	defer tty.Close()
	return term.ReadPassword(int(tty.Fd()))
}

func readPasswordFromWindows(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	return term.ReadPassword(int(os.Stdin.Fd()))
}
