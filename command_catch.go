package main

import (
	"fmt"
	"math/rand"
)

// catchThreshold sets the catch probability to threshold / base_experience,
// so Pokemon with higher base experience are harder to catch.
const catchThreshold = 40

func commandCatch(cfg *config, s []string) error {
	if len(s) == 0 {
		fmt.Println("Catch requires a Pokemon name")
		return nil
	}

	name := s[0]
	fmt.Println("Throwing a Pokeball at " + name + "...")

	pokemon, err := cfg.client.GetPokemon(name)
	if err != nil {
		return err
	}

	caught := pokemon.BaseExperience <= catchThreshold || rand.Intn(pokemon.BaseExperience) < catchThreshold
	if !caught {
		fmt.Println(name + " escaped!")
		return nil
	}

	fmt.Println(name + " was caught!")
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}
