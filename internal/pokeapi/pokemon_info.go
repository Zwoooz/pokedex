package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(target string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + target

	var info Pokemon
	if cached, ok := c.cache.Get(target); ok {
		err := json.Unmarshal(cached, &info)
		if err != nil {
			return Pokemon{}, err
		}
		return info, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("%s does not exist", target)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(target, data)

	err = json.Unmarshal(data, &info)
	if err != nil {
		return Pokemon{}, err
	}
	return info, nil
}
