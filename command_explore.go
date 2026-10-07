package main

import (
	"fmt"
)

func commandExplore(cfg *config, s []string) error {
	if len(s) == 0 {
		fmt.Println("Explore requires an area name")
		return nil
	}

	areaName := s[0]
	fmt.Println("Exploring " + areaName + "...")
	result, err := cfg.client.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
