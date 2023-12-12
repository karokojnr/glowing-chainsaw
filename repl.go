package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		switch command {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Welcome to Pokedex help menu!")
			fmt.Println("Available commands:")
			fmt.Println("exit - exit the program")
			fmt.Println("help - show this help menu")
			fmt.Println("")
		default:
			fmt.Println("Invalid command")
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
