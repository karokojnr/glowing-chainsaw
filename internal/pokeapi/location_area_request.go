package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {

	fullURL := baseURL + "location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	// * check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// * found in cache
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}

	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// * saved to cache
	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}

func (c *Client) GetListLocationArea(locationAreaName string) (LocationArea, error) {

	fullURL := baseURL + "location-area/" + locationAreaName

	// * check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// * found in cache
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// * saved to cache
	c.cache.Add(fullURL, dat)

	return locationArea, nil

}
