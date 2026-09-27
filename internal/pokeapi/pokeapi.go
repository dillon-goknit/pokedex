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

const BaseURL = "https://pokeapi.co/api/v2/location-area"

func GetLocationAreas(url string) (LocationAreasResp, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("status %d: %s", res.StatusCode, body)
	}

	var data LocationAreasResp
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreasResp{}, err
	}
	return data, nil
}
