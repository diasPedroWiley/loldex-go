package dddragon

import (
	"encoding/json"
	"net/http"
)

const baseUrl = "http://ddragon.leagueoflegends.com"

func Handle(urlParams string) ([]string, error) {
	finalUrl := baseUrl + urlParams
	var response []string

	resp, err := http.Get(finalUrl)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}
