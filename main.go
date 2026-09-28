package main

import (
	"time"

	"github.com/dillon-goknit/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		client:   pokeapi.NewClient(5 * time.Minute),
	}
	startRepl(cfg)
}
