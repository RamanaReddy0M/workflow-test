package main

import (
	"fmt"
	"os"

	"github.com/projectdiscovery/gologger"
	"golang.org/x/term"
	// Assuming gologger is properly imported.
)

func main() {
	fmt.Println("int(os.Stdin.Fd()): ", int(os.Stdin.Fd()))
	fmt.Printf("[*] Enter PDCP API Key (type 'exit' to abort): ")

	// Open /dev/tty for direct terminal input, bypassing stdin.
	tty, err := os.Open("/dev/tty")
	if err != nil {
		gologger.Fatal().Msgf("Could not open /dev/tty: %s\n", err)
	}
	defer tty.Close()

	// Use term.ReadPassword to securely read the input from the terminal.
	apiKeyBytes, err := term.ReadPassword(int(tty.Fd()))
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
