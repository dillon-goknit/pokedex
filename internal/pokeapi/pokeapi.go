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
