package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreaResp, error) {
	endPoint := "/location-area"
	fullURL := baseURL + endPoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)

	if err != nil {
		return locationAreaResp, err
	}

	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endPoint := baseURL + "/location-area/" + locationAreaName

	req, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil

}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endPoint := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil

}
