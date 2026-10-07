package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dillon-goknit/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commands map[string]cliCommand
	client   *pokeapi.Client
	next     *string
	previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all the Pokemon in a given  area",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

func commandResponse(command string) string {
	words := cleanInput(command)
	if len(words) == 0 {
		return ""
	}
	first := words[0]
	return "Your command was: " + first
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandRest := words[1:]
		command, ok := cfg.commands[commandName]
		if !ok {
			fmt.Println("unknown command")
			continue
		}

		err := command.callback(cfg, commandRest)
		if err != nil {
			fmt.Println(err)
		}
	}
}
