package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex: ")

	for _, pokemon := range cfg.caughPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
