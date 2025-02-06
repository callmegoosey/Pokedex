package pokeapi

import (
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

//var poke_LocationArea PokeLocationAreas

func CommandMap(url string) (array_of_bytes []byte, a error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// if err := json.Unmarshal(body, &poke_LocationArea); err != nil {
	// 	return nil, err
	// }

	// //print all 20 location
	// for _, location := range poke_LocationArea.Results {
	// 	fmt.Println(location.Name)
	// }

	// return poke_LocationArea.Next, poke_LocationArea.Previous, nil
	return body, nil
}
