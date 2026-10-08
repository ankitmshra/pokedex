package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL         = "https://pokeapi.co"
	apiVersion      = "/api/v2"
	apiBaseURL      = baseURL + apiVersion
	locationAreaURL = apiBaseURL + "/location-area/"
	pokemonURL      = apiBaseURL + "/pokemon/"
)

func getDataFromAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
