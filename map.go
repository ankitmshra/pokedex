package main

import (
	"errors"
	"fmt"
)

type navType string

const (
	Next     navType = "Next"
	Previous navType = "Previous"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (la *locationArea) getUrl(nt navType) *string {
	var url string
	if la == nil {
		url = locationAreaURL
		return &url
	}

	switch nt {
	case Next:
		if la.Next != nil {
			url = *la.Next
		} else {
			fmt.Println("you're on the last page")
			return nil
		}
	case Previous:
		if la.Previous != nil {
			url = *la.Previous
		} else {
			fmt.Println("you're on the first page")
			return nil
		}
	default:
		return nil
	}
	return &url
}

func getMapData(cnf *config, nt navType) ([]byte, error) {
	la := cnf.locationArea
	url := la.getUrl(nt)
	if url == nil {
		return nil, errors.New("no location area found")
	}

	var data []byte
	// try to receive from cache
	if cache := cnf.cache; cache != nil {
		data, ok := cache.Get(*url)
		// in case of cache hit
		if ok {
			fmt.Println("Cache Hit")
			return data, nil
		}
	}

	// in case of cache miss
	// Fetch the data from API
	data, err := getDataFromAPI(*url)

	if err != nil {
		return nil, err
	}

	// add back to the cache
	cnf.cache.Add(*url, data)
	return data, nil

}

type LocationData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getLocationData(cnf *config, location string) ([]byte, error) {
	url := locationAreaURL + location + "/"

	if cache := cnf.cache; cache != nil {
		data, ok := cache.Get(url)
		// in case of cache hit
		if ok {
			fmt.Println("Cache Hit")
			return data, nil
		}
	}

	// in case of cache miss
	// Fetch the data from API
	data, err := getDataFromAPI(url)
	if err != nil {
		fmt.Println("Unable to find location data")
		return nil, err
	}
	cnf.cache.Add(url, data)
	return data, nil
}
