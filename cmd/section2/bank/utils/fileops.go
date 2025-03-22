package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetValueFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("failed to find file")
	}

	valueStr := string(data)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, errors.New("failed to parse value into float")
	}

	return value, nil
}

func WriteValueToFile(filename string, value float64) {
	valueStr := fmt.Sprint(value)

	os.WriteFile(filename, []byte(valueStr), 0644)
}
