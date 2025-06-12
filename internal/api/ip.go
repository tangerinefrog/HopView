package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type locationResponse struct {
	Location struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"location"`
}

const ipProviderUrl string = "https://api.ipgeolocation.io/v2/ipgeo?apiKey=%s&ip=%s&fields=location.latitude,location.longitude"

func GetIpLocation(ip string) (locationResponse, error) {
	apiKey := os.Getenv("IP_API_KEY")

	url := fmt.Sprintf(ipProviderUrl, apiKey, ip)
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
