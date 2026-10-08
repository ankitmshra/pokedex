package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, args ...string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

map: next lsit of location areas
mapb: previous list of location areas
help: Displays a help message
exit: Exit the Pokedex`)

	return nil
}

func commandMap(cnf *config, args ...string) error {
	data, err := getMapData(cnf, Next)
	if err != nil {
		return err
	}

	cnf.locationArea = &locationArea{}

	if err = json.Unmarshal(data, cnf.locationArea); err != nil {
		return err
	}

	for _, result := range cnf.locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(cnf *config, args ...string) error {
	data, err := getMapData(cnf, Previous)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, cnf.locationArea); err != nil {
		return err
	}

	for _, result := range cnf.locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExplore(cnf *config, args ...string) error {
	cnf.locationData = &LocationData{}
	data, err := getLocationData(cnf, args[0])
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, cnf.locationData); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, results := range cnf.locationData.PokemonEncounters {
		fmt.Println(results.Pokemon.Name)
	}

	return nil
}

func commandCatch(cnf *config, args ...string) error {
	ld := cnf.locationData
	pokemonName := args[0]
	if ld == nil {
		return fmt.Errorf("explore the map to find a pokemon")
	}

	if _, ok := cnf.pokemon[pokemonName]; ok {
		return fmt.Errorf("%s already caught", pokemonName)
	}

	pokemonFound := false
	for _, pe := range ld.PokemonEncounters {
		if pe.Pokemon.Name == pokemonName {
			pokemonFound = true
		}
	}

	if !pokemonFound {
		return fmt.Errorf("pokemon not found in this location")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	caught, err := catchPokemon(cnf, pokemonName)
	if err != nil {
		return err
	}
	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(cnf *config, args ...string) error {
	pokemonName := args[0]
	pokemon, ok := cnf.pokemon[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon not caught yet")
	}

	pokemon.printStats()
	return nil
}

func commandPokedex(cnf *config, args ...string) error {
	if cnf.pokemon == nil {
		return fmt.Errorf("no pokemon caught yet")
	}

	for pokemon, _ := range cnf.pokemon {
		fmt.Println(pokemon)
	}

	return nil
}
