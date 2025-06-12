package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationResponse struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

const ipProviderUrl string = "http://ip-api.com/json/%s?fields=lat,lon"

func GetIpLocation(ip string) (locationResponse, error) {
	url := fmt.Sprintf(ipProviderUrl, ip)
	res, err := http.Get(url)
	if err != nil {
		return locationResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return locationResponse{}, err
		}

		body := locationResponse{}
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			return locationResponse{}, err
		}

		return body, nil
	}

	return locationResponse{}, fmt.Errorf("got unsuccessful status code from '%s'", url)
}
