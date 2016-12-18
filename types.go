package mtg

import (
	"encoding/json"
	"net/http"
)

// GetTypes fetches a list of all card types
func GetTypes() ([]string, error) {
	resp, err := http.Get(queryUrl + "types")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, err
	}
	res := new(struct {
		Types []string `json:"types"`
	})
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}
	return res.Types, nil
}

// GetSuperTypes fetches a list of all card supertypes
func GetSuperTypes() ([]string, error) {
	resp, err := http.Get(queryUrl + "supertypes")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, err
	}
	res := new(struct {
		Types []string `json:"supertypes"`
	})
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}
	return res.Types, nil
}

// GetSubTypes fetches a list of all card subtypes
func GetSubTypes() ([]string, error) {
	resp, err := http.Get(queryUrl + "subtypes")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, err
	}
	res := new(struct {
		Types []string `json:"subtypes"`
	})
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}
	return res.Types, nil
}

// GetFormats fetches a list of all known game formats
func GetFormats() ([]string, error) {
	resp, err := http.Get(queryUrl + "formats")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, err
	}
	res := new(struct {
		Formats []string `json:"formats"`
	})
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}
	return res.Formats, nil
}
