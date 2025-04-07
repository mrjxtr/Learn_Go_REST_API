package conversion

import (
	"errors"
	"strconv"
)

func StrsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for idx, strVal := range strings {
		floatPrice, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, errors.New("failes to parse string to float")
		}

		floats[idx] = floatPrice
	}
	return floats, nil
}
