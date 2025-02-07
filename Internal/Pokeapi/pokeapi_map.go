package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeLocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeLocation struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// submits a http request and gets the body
func Get(url string) (array_of_bytes []byte, a error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// ===========================================
//
//	Returns 20 map name
//
// ===========================================
func Get_map_names(array_of_bytes *[]byte) (Next string, Prev string, err error) {
	poke_location_areas := PokeLocationAreas{}

	if err := json.Unmarshal(*array_of_bytes, &poke_location_areas); err != nil {
		return "", "", err
	}

	//print all 20 location
	for _, location := range poke_location_areas.Results {
		fmt.Println(location.Name)
	}

	return poke_location_areas.Next, poke_location_areas.Previous, nil
}

// ===========================================
//
//	Given a valid map name
//	returns all pokemon in the area
//
// ===========================================
func Explore_map(array_of_bytes *[]byte, other_inputs *[]string) error {
	poke_location_areas := PokeLocation{}

	if err := json.Unmarshal(*array_of_bytes, &poke_location_areas); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", (*other_inputs)[0])
	fmt.Printf("Found Pokemon:\n")

	for _, pokemon := range poke_location_areas.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
