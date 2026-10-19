package main

import (
	"fmt"
)

func commandInspect(cfg *config, s []string) error {
	if len(s) == 0 {
		fmt.Println("Inspect requires a Pokemon name")
		return nil
	}

	name := s[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Println("  -", pokemonType.Type.Name)
	}
	return nil
}
