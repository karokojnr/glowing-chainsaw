package main

import (
	"time"

	"github.com/karokojnr/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughPokemon            map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughPokemon:  make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
