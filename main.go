package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const banner = `
  ███████╗██╗     ██╗███████╗ █████╗ 
  ██╔════╝██║     ██║╚══███╔╝██╔══██╗
  █████╗  ██║     ██║  ███╔╝ ███████║
  ██╔══╝  ██║     ██║ ███╔╝  ██╔══██║
  ███████╗███████╗██║███████╗██║  ██║
  ╚══════╝╚══════╝╚═╝╚══════╝╚═╝  ╚═╝

  ELIZA — DOCTOR script (Weizenbaum, 1966)
  Type your message and press Enter.
  Type "quit" or Ctrl-D to exit.
`

func main() {
	once := flag.String("once", "", "send a single message and print the response, then exit")
	flag.Parse()

	e := New()

	// --once mode: non-interactive, useful for piped input or scripting.
	if *once != "" {
		fmt.Println(e.Respond(*once))
		return
	}

	fmt.Print(banner)
	fmt.Println("ELIZA: How do you do. Please state your problem.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("You:   ")
		if !scanner.Scan() {
			// EOF (Ctrl-D)
			fmt.Println("\nGoodbye.")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		response := e.Respond(input)
		fmt.Printf("ELIZA: %s\n\n", response)

		// Quit after printing the farewell.
		lower := strings.ToLower(input)
		if lower == "quit" || lower == "bye" || lower == "goodbye" || lower == "exit" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
}
