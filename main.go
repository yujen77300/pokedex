package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	fmt.Println("Help command executed")
}

func commandExit() {
	fmt.Println("Exit command executed")
}

func main() {
	commandMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	onCli := true
	scanner := bufio.NewScanner(os.Stdin)

	for onCli {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if input == "help" {
			fmt.Println(`
Welcome to the Pokedex!
Usage:`)

			fmt.Println("")

			for n, d := range commandMap {
				fmt.Println(n + ":  " + d.description)
			}

			fmt.Println("")
		} else if input == "exit" {
			onCli = false
		}

	}

}
