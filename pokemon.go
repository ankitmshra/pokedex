package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"pokedex/internal/pokecache"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (p *Pokemon) printStats() {
	fmt.Println("Name: ", p.Name)
	fmt.Println("Height: ", p.Height)
	fmt.Println("Weight: ", p.Weight)
	fmt.Println("Stats: ")
	for _, stat := range p.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, typ := range p.Types {
		fmt.Printf("\t-%s\n", typ.Type.Name)
	}
}

func caught(baseExperience int) bool {
	generatedNumber := rand.IntN(100)
	catchChance := 90

	switch {
	case baseExperience > 90:
		catchChance = 10
	case baseExperience > 60:
		catchChance = 30
	case baseExperience > 30:
		catchChance = 60
	}

	return generatedNumber < catchChance

}

func getPokemonData(ca *pokecache.Cache, url string) ([]byte, error) {
	// try to receive from cache
	if ca != nil {
		data, ok := ca.Get(url)
		// in case of cache hit
		if ok {
			fmt.Println("Cache Hit")
			return data, nil
		}
	}

	data, err := getDataFromAPI(url)
	if err != nil {
		return nil, err
	}

	if ca != nil {
		ca.Add(url, data)
	}
	return data, nil
}

func catchPokemon(cnf *config, pokemonName string) (bool, error) {
	url := pokemonURL + pokemonName + "/"

	data, err := getPokemonData(cnf.cache, url)
	if err != nil {
		return false, err
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return false, err
	}

	if caught := caught(pokemon.BaseExperience); caught {
		cnf.pokemon[pokemonName] = pokemon
		return true, nil
	}

	return false, nil
}
