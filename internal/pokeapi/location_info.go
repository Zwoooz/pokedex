package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationInfo(locationName string) (pokeSingleLocation, error) {
	fullURL := baseURL + "/location-area/" + locationName

	var location pokeSingleLocation
	if cached, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return pokeSingleLocation{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return pokeSingleLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokeSingleLocation{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return pokeSingleLocation{}, fmt.Errorf("%s does not exist", locationName)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeSingleLocation{}, err
	}
	c.cache.Add(fullURL, data)

	err = json.Unmarshal(data, &location)
	if err != nil {
		return pokeSingleLocation{}, err
	}
	return location, nil
}
