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
	data, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreasResp{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return LocationAreasResp{}, fmt.Errorf("unexpected status: %s", res.Status)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasResp{}, err
		}
		c.cache.Add(url, data)
	}

	var resp LocationAreasResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationAreasResp{}, err
	}
	return resp, nil
}

func (c *Client) GetLocationArea(areaName string) (LocationAreaDetail, error) {
	url := fmt.Sprintf("%s/%s", BaseURL, areaName)
	data, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaDetail{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return LocationAreaDetail{}, fmt.Errorf("unexpected status: %s", res.Status)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetail{}, err
		}

		c.cache.Add(url, data)
	}

	var resp LocationAreaDetail
	if err := json.Unmarshal(data, &resp); err != nil {
		return LocationAreaDetail{}, err
	}

	return resp, nil
}
