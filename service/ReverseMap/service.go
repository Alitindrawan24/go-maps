package reversemap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://nominatim.openstreetmap.org"
)

func Execute(lat string, lon string) (ReverseMaps, error) {
	var client = &http.Client{}
	var reverseMap ReverseMaps

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/reverse?lat=%s&lon=%s&zoom=18&format=json", baseURL, lat, lon), nil)
	if err != nil {
		return reverseMap, err
	}

	response, err := client.Do(request)
	if err != nil {
		return reverseMap, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&reverseMap)
	if err != nil {
		return reverseMap, err
	}

	return reverseMap, nil
}
