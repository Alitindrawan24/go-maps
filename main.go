package go_maps

import reversemap "github.com/Alitindrawan24/go-maps/service/ReverseMap"

func ReverseMaps(lat string, lon string) (reversemap.ReverseMaps, error) {
	reverseMap, err := reversemap.Execute(lat, lon)

	if err != nil {
		return reverseMap, err
	}

	return reverseMap, nil
}
