package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (pokeLocations, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var locations pokeLocations
	if cached, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(cached, &locations)
		if err != nil {
			return pokeLocations{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return pokeLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokeLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeLocations{}, nil
	}
	c.cache.Add(fullURL, data)

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return pokeLocations{}, nil
	}
	return locations, nil
}
