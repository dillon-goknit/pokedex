package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResp struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
}

type LocationAreaDetail struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

const BaseURL = "https://pokeapi.co/api/v2/location-area"

func (c *Client) GetLocationAreas(url string) (LocationAreasResp, error) {
	data, err := c.fetch(url)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var resp LocationAreasResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationAreasResp{}, err
	}
	return resp, nil
}

func (c *Client) GetLocationArea(areaName string) (LocationAreaDetail, error) {
	url := fmt.Sprintf("%s/%s", BaseURL, areaName)
	data, err := c.fetch(url)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	var resp LocationAreaDetail
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationAreaDetail{}, err
	}
	return resp, nil
}

func (c *Client) fetch(url string) ([]byte, error) {
	if data, ok := c.cache.Get(url); ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, data)
	return data, nil
}
