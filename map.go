package main

import (
	"fmt"

	"github.com/dillon-goknit/pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := pokeapi.BaseURL
	if cfg.next != nil {
		url = *cfg.next
	}
	return showLocationAreas(cfg, url)
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	return showLocationAreas(cfg, *cfg.previous)
}

func showLocationAreas(cfg *config, url string) error {
	resp, err := cfg.client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.next = resp.Next
	cfg.previous = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
