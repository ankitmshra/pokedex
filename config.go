package main

import (
	"pokedex/internal/pokecache"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
type config struct {
	commandRegistry map[string]cliCommand
	locationArea    *locationArea
	locationData    *LocationData
	cache           *pokecache.Cache
	pokemon         map[string]Pokemon
}

func (c *config) initConfig() {
	*c = config{
		commandRegistry: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit this pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Show this help",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Navigate pokedex locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Navigate back pokedex locations",
				callback:    commandMapB,
			},
			"explore": {
				name:        "explore",
				description: "Navigate pokedex locations",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Catch available pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect pokemon",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Show all the caught pokemons",
				callback:    commandPokedex,
			},
		},
		locationArea: nil,
		locationData: nil,
		cache:        pokecache.NewCache(5 * time.Minute),
		pokemon:      make(map[string]Pokemon),
	}
}
