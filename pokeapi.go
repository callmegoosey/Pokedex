package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeLocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// type pokeLocation struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Region struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"region"`
// 	Names []struct {
// 		Name     string `json:"name"`
// 		Language struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"language"`
// 	} `json:"names"`
// 	GameIndices []struct {
// 		GameIndex  int `json:"game_index"`
// 		Generation struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"generation"`
// 	} `json:"game_indices"`
// 	Areas []struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"areas"`
// }

var poke_LocationArea pokeLocationAreas

func commandMap(url string, ptr_config *config) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &poke_LocationArea); err != nil {
		return err
	}

	//print all 20 location
	for _, location := range poke_LocationArea.Results {
		fmt.Println(location.Name)
	}

	//update next and previous
	ptr_config.next = &poke_LocationArea.Next
	ptr_config.previous = &poke_LocationArea.Previous
	return nil
}

func commandMap_forward(ptr_config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.next != nil {
		url = *ptr_config.next
	}

	return commandMap(url, ptr_config)
}

func commandMap_backward(ptr_config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.previous != nil {
		url = *ptr_config.previous
	}

	return commandMap(url, ptr_config)
}
