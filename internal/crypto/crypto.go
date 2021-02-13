package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPairPrice(pair string) (float64, error) {
	res, err := http.Get(fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%v", pair))
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil
	}

	type krakenResponse struct {
		Error  []string
		Result map[string]struct {
			Ask [3]string `json:"a"`
		} `json:"result"`
	}

	var response krakenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	if len(response.Error) != 0 {
		return 0, errors.New(response.Error[0])
	}

	var price float64

	for key := range response.Result {
		value, _ := response.Result[key]
		price, err = strconv.ParseFloat(value.Ask[0], 64)
		if err != nil {
			return 0, err
		}
	}

	return price, nil
}

func GetPairList() ([]string, error) {
	res, err := http.Get("https://api.kraken.com/0/public/AssetPairs")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Result map[string]interface{}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	pairs := make([]string, 0)
	for pair := range response.Result {
		pairs = append(pairs, pair)
	}

	return pairs, nil
}
